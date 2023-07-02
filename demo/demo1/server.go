package main

import (
	"fmt"
	"io"
	"net"
	"sync"
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
		}
	}()

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
