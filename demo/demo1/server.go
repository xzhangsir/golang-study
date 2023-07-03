package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int
	// 在线用户的列表
	OnLineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的channel
	Message chan string
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnLineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 监听Message广播消息channel的goroutine
// 一旦有消息就发送给全部的在线user
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message
		// 将msg发送给全部的在线user
		this.mapLock.Lock()
		for _, user := range this.OnLineMap {
			user.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	this.Message <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	// fmt.Println("链接成功")
	// 用户上线了 将用户加入到OnlineMap中
	user := NewUser(conn, this)
	user.Online()
	// 监听用户是否是活跃的channel
	isLive := make(chan bool)
	// 接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				// 用户下线
				user.OffLine()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn read err", err)
				return
			}
			// 读取用户的消息 去除\n
			msg := string(buf[:n-1])
			fmt.Println(msg)
			// 将消息进行广播
			user.DoMessage(msg)
			// 用户的任意消息 代表当前的用户是活跃的
			isLive <- true
		}
	}()

	for{
		select{
		case <- isLive:
			// 当前用户活跃 重置定时器
			// 不做任何事情 为了激活select 更新下面的定时器
		case <- time.After(this.Second * 10):
			// 已经超时 强制退出
			user.SendMsg("你被踢了")
			// 销毁用户资源
			close(user.C)
			// 关闭链接
			conn.Close()
			return
		}
	}

}

// 启动服务的方法
func (this *Server) Start() {
	// net.Listen("tcp", "127.0.0.1:8888")
	Listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listener", err)
		return
	}
	defer Listener.Close()
	// 启动监听message的goroutine
	go this.ListenMessager()
	for {
		conn, err := Listener.Accept()
		if err != nil {
			fmt.Println("Listner accept err", err)
			continue
		}
		go this.Handler(conn)
	}

}
