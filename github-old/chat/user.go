package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
	// 当前用户属于哪个server
	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	go user.ListenMessage()
	return user
}

// 用户上线功能
func (this *User) Online() {
	this.server.mapLock.Lock()
	// 用户上线了 将用户加入到OnlineMap中
	this.server.OnLineMap[this.Name] = this
	this.server.mapLock.Unlock()
	// 广播当前用户上线消息
	this.server.BroadCast(this, "已上线了")
}

// 用户下线功能
func (this *User) OffLine() {
	this.server.mapLock.Lock()
	// 用户下线了 将用户从OnlineMap中删除
	delete(this.server.OnLineMap, this.Name)
	this.server.mapLock.Unlock()
	// 广播当前用户下线消息
	this.server.BroadCast(this, "下线")
}

// 给当前user对应的客户端发消息
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 用户处理消息的业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		// 查询当前在线用户都有哪些
		this.server.mapLock.Lock()
		for _, user := range this.server.OnLineMap {
			onLineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线..\n"
			this.SendMsg(onLineMsg)
		}
		this.server.mapLock.Unlock()

	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 修改用户名 rename|zhangsan
		newName := msg[7:]
		// 判断name是否存在
		_, ok := this.server.OnLineMap[newName]
		if ok {
			this.SendMsg("当前用户名存在")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnLineMap, this.Name)
			this.server.OnLineMap[newName] = this
			this.server.mapLock.Unlock()
			this.SendMsg("修改用户名称成功")
		}
		this.Name = msg[7:]
	} else if len(msg) > 4 && msg[:3] == "to|"{
		// 私聊消息格式 to|zhangsan|消息内容
		username := strings.Split(msg,"|")[1]
		if username == ""{
			this.SendMsg("消息格式错误")
			return
		}
		user,ok := this.server.OnLineMap[username]
		if !ok {
			this.SendMsg("当前用户不存在")
			return
		} else {
			user.SendMsg(this.Name + "对您说：" + strings.Split(msg,"|")[2])
		}

	}else {
		this.server.BroadCast(this, msg)
	}
}

// 监听当前User channel的方法，
// 一旦有消息就直接发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		fmt.Println(msg)
		this.conn.Write([]byte(msg + "\n"))
	}
}
