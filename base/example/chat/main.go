package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	ip   string
	port int
	// 在线用户
	onlineUser map[string]*User
	mapLock    sync.RWMutex
	// 消息广播的channel
	message chan string
}

func NewServer(ip string, port int) *Server {
	return &Server{
		ip:         ip,
		port:       port,
		onlineUser: make(map[string]*User),
		message:    make(chan string),
	}
}

func (_this *Server) ListenMessager() {
	for {
		msg := <-_this.message
		// 将msg发送给全部的在线user
		_this.mapLock.Lock()
		for _, user := range _this.onlineUser {
			user.C <- msg
		}
		_this.mapLock.Unlock()
	}
}

func (_this *Server) Handler(conn net.Conn) {
	user := NewUser(conn, _this)
	user.Online()
	// 监听用户是否是活跃的channel
	isLive := make(chan bool)
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

	for {
		select {
		case <-isLive:
			// 当前用户活跃 重置定时器
			// 不做任何事情 为了激活select 更新下面的定时器
		case <-time.After(time.Second * 1000):
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

// 广播消息
func (_this *Server) BroadCast(user *User, message string) {
	sendMsg := "[" + user.addr + "]" + user.name + ":" + message
	_this.message <- sendMsg
}

// 启动server服务
func (_this *Server) Start() {
	// 监听
	Listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", _this.ip, _this.port))
	if err != nil {
		fmt.Println("net.Listener", err)
		return
	}
	defer Listener.Close()
	go _this.ListenMessager()
	for {
		// 建立连接
		conn, err := Listener.Accept()
		if err != nil {
			fmt.Println("Listner accept err", err)
			continue
		}
		go _this.Handler(conn)
	}

}

func main() {
	s := NewServer("127.0.0.1", 8001)
	s.Start()
}
