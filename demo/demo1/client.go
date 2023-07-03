package main

import (
	"net"
	"fmt"
	"flag"
	"io"
	"os"
)

type Client struct{
	ServerIp string
	ServerPort int
	Name string
	conn net.Conn
	flag int //当前客户端的模式
}

func NewClient(serverIp string,serverPort int) *Client{
	// 创建客户端
	client := &Client{
		ServerIp:serverIp,
	  ServerPort:serverPort,
		flag:999,
	}
	// 链接server
	conn,err := net.Dial("tcp",fmt.Sprintf("%s:%d",serverIp,serverPort))
	if err != nil{
		fmt.Println("netDial err",err)
		return 
	}
	client.conn = conn

	// 返回对象
	return client
}

// 接受服务器的响应
func (client *Client) DealRes(){
	// 一旦client.conn 有数据 就直接copy到stdout标准输出上
	// 永久监听
 io.Copy(os.Stdout,client.conn)
}


func (client *Client) menu() bool{
		var flag int
		fmt.Println("1.公聊模式")
		fmt.Println("2.私聊模式")
		fmt.Println("3.更新用户名")
		fmt.Println("0.退出")
		fmt.Scanln(&flag)
		if flag >= 0 && flag <= 3{
			client.flag = flag
			return true
		}else{
			fmt.Println("输入不合法")
			return false
		}
}

func (client *Client) UpdateName() bool{
	fmt.Println("请输入用户名")
	fmt.Scanln(&client.Name)
	sendMsg := "rename|" + client.Name + "\n"
	_,err := client.conn.Write([]byte(sendMsg))
	if err != nil{
		fmt.Println("conn.Write err",err)
		return false
	}
	return true
}

func (client *Client) PublicChat(){
	var chatMsg string
  fmt.Println("请输入消息，exit退出")
	fmt.Scanln(&chatMsg)
	for chatMsg != "exit"{
		// 发送给服务器
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_,err := client.conn.Write([]byte(sendMsg))
			if err != nil{
				fmt.Println("conn.Write err",err)
				break
			}
		}
		chatMsg = ""
		fmt.Println("请输入消息，exit退出")
		fmt.Scanln(&chatMsg)
	}
}

// 查询在线用户
func (client *Client) SelectUsers(){
	sendMsg := "who\n"
	_,err := client.conn.Write([]byte(sendMsg))
	if err != nil{
		fmt.Println("conn.Write err",err)
		return
	}
}

func (client *Client) PrivateChat(){
	client.SelectUsers()
	var userName string
	
  fmt.Println("请输入聊天对象的用户名 ，exit退出")
  fmt.Scanln(&userName)
	for userName != "exit"{
		fmt.Println("请输入消息，exit退出")
		var chatMsg string
		fmt.Scanln(&chatMsg)
		for chatMsg != "exit"{
			if len(chatMsg) != 0 {
				sendMsg := "to|" + userName + "|" +chatMsg + "\n"
				_,err := client.conn.Write([]byte(sendMsg))
				if err != nil{
					fmt.Println("conn.Write err",err)
					break
				}
			}
			chatMsg = ""
			fmt.Println("请输入消息，exit退出")
			fmt.Scanln(&chatMsg)
		}
	}
}

func (client *Client) Run(){
	for client.flag != 0 {
		for client.menu() != true{
		}
		// 根据不同模式 处理不同业务
		switch client.flag {
		case 1:
			fmt.Println("选择了公聊模式")
			client.PublicChat()
			break
		case 2:
			fmt.Println("选择了私聊模式")
			client.PrivateChat()
			break
		case 3:
			fmt.Println("选择了更新用户名")
			client.UpdateName()
			break
		}
	}
}



var serverIp string 
var serverPort int

// client -ip 127.0.0.1 -port 8888
func init(){
	flag.StringVar(&serverIp,"ip","127.0.0.1","设置服务器IP地址")
	flag.IntVar(&serverPort,"port",8888,"设置服务器端口")
}

func main(){
	// 命令行解析
	flag.Parse()
	client := NewClient(serverIp,serverPort)
	// client := NewClient("127.0.0.1",8888)
	if client == nil{
		fmt.Println(">>>链接服务器失败")
		return 
	}
	
	go client.DealRes()
	fmt.Println(">>>链接服务器成功")
	// 启动
	client.Run()
}