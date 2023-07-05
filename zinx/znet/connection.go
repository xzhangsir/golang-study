package znet

import (
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
	//该连接的处理方法api
	handleAPI ziface.HandFunc
	//告知该链接已经退出/停止的channel
	ExitBuffChan chan bool
}

// 创建连接
func NewConntion(conn *net.TCPConn, connId uint32, callback_api ziface.HandFunc) *Connection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connId,
		isClosed:     false,
		handleAPI:    callback_api,
		ExitBuffChan: make(chan bool, 1),
	}
	return c
}

// 启动连接，让当前连接开始工作
func (c *Connection) Start() {

}
func (c *Connection) Stop() {

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
