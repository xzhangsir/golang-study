package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	name string
	addr string
	C    chan string
	conn net.Conn
	// 当前用户属于哪个server
	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	// 获取远端网络地址
	userAddr := conn.RemoteAddr().String()
	user := &User{
		name:   userAddr,
		addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	go user.ListenMessage()
	return user
}

// 用户上线功能
func (_this *User) Online() {
	_this.server.mapLock.Lock()
	// 用户上线了 将用户加入到onlineUser中
	_this.server.onlineUser[_this.name] = _this
	_this.server.mapLock.Unlock()
	// 广播当前用户上线消息
	_this.server.BroadCast(_this, "来新人了")
}

// 用户下线功能
func (_this *User) OffLine() {
	_this.server.mapLock.Lock()
	// 用户下线了 将用户从onlineUser中删除
	delete(_this.server.onlineUser, _this.name)
	_this.server.mapLock.Unlock()
	// 广播当前用户下线消息
	_this.server.BroadCast(_this, "退下")
}

// 给当前user对应的客户端发消息
func (_this *User) SendMsg(msg string) {
	_this.conn.Write([]byte(msg))
}

// 用户处理消息的业务
func (_this *User) DoMessage(msg string) {
	if msg == "who" {
		// 查询当前在线用户都有哪些
		_this.server.mapLock.Lock()
		for _, user := range _this.server.onlineUser {
			onLineMsg := "[" + user.addr + "]" + user.name + ":" + "在线..\n"
			_this.SendMsg(onLineMsg)
		}
		_this.server.mapLock.Unlock()

	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 修改用户名 rename|zhangsan
		newName := msg[7:]
		// 判断name是否存在
		_, ok := _this.server.onlineUser[newName]
		if ok {
			_this.SendMsg("当前用户名存在")
		} else {
			_this.server.mapLock.Lock()
			delete(_this.server.onlineUser, _this.name)
			_this.server.onlineUser[newName] = _this
			_this.server.mapLock.Unlock()
			_this.SendMsg("修改用户名称成功")
		}
		_this.name = msg[7:]
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 私聊消息格式 to|zhangsan|消息内容
		username := strings.Split(msg, "|")[1]
		if username == "" {
			_this.SendMsg("消息格式错误")
			return
		}
		user, ok := _this.server.onlineUser[username]
		if !ok {
			_this.SendMsg("当前用户不存在")
			return
		} else {
			user.SendMsg(_this.name + "对您说：" + strings.Split(msg, "|")[2])
		}

	} else {
		_this.server.BroadCast(_this, msg)
	}
}

// 监听当前User channel的方法，
// 一旦有消息就直接发送给客户端
func (_this *User) ListenMessage() {
	for {
		msg := <-_this.C
		fmt.Println(msg)
		_this.conn.Write([]byte(msg + "\n"))
	}
}
