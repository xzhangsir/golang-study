package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/ziface"
)

type Connection struct {
	//当前连接的socket TCP套接字
	Conn *net.TCPConn
	//当前连接的ID 也可以称作为SessionID，ID全局唯一
	ConnID uint32
	//当前连接是否关闭
	isClosed bool
	//该连接的处理方法router
	Router ziface.IRouter
	// //该连接的处理方法api
	// handleAPI ziface.HandFunc
	//告知该链接已经退出/停止的channel
	ExitBuffChan chan bool
}

// 创建连接
func NewConntion(conn *net.TCPConn, connId uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connId,
		isClosed: false,
		// handleAPI:    callback_api,
		Router:       router,
		ExitBuffChan: make(chan bool, 1),
	}
	return c
}

// 启动连接，让当前连接开始工作
func (c *Connection) Start() {
	// 开启处理该链接 读取到客户端数据之后的业务请求
	go c.StartReader()
	// for {
	// 	select {
	// 	case <-c.ExitBuffChan:
	// 		// 得到退出消息 不再阻塞
	// 		return
	// 	}
	// }
	// 优化如下
	for range c.ExitBuffChan {
		return
	}

}
func (c *Connection) Stop() {
	// 如果当前链接已经关闭
	if c.isClosed {
		return
	}
	c.isClosed = true
	// 关闭socket链接
	c.Conn.Close()
	//通知从缓冲队列读数据的业务，该链接已经关闭
	c.ExitBuffChan <- true
	//关闭该链接全部管道
	close(c.ExitBuffChan)

}
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 处理conn读数据的Goroutine
func (c *Connection) StartReader() {
	fmt.Println("处理conn读数据的goroutine在运行")
	defer fmt.Println(c.RemoteAddr().String(), "conn exit")
	defer c.Stop()
	for {
		// 创建拆包解包的对象
		dp := NewDataPack()
		//读取客户端的Msg head
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error ", err)
			c.ExitBuffChan <- true
			continue
		}
		//拆包，得到msgid 和 datalen 放在msg中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error ", err)
			c.ExitBuffChan <- true
			continue
		}
		//根据 dataLen 读取 data，放在msg.Data中
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error ", err)
				c.ExitBuffChan <- true
				continue
			}
		}
		msg.SetData(data)
		//得到当前客户端请求的Request数据
		req := Request{
			conn: c,
			msg:  msg, //将之前的buf 改成 msg
		}
		//从路由Routers 中找到注册绑定Conn的对应Handle
		go func(request ziface.IRequest) {
			//执行注册的路由方法
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

// 直接将Message数据发送数据给远程的TCP客户端
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("Connection closed when send msg")
	}
	//将data封包，并且发送
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId)
		return errors.New("Pack error msg ")
	}

	//写回客户端
	if _, err := c.Conn.Write(msg); err != nil {
		fmt.Println("Write msg id ", msgId, " error ")
		c.ExitBuffChan <- true
		return errors.New("conn Write error")
	}

	return nil
}
