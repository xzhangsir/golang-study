package main;

import (
	"net"
	"fmt"
)

func main(){
	// 1 本地端口启动服务
 	listener,err := net.Listen("tcp","127.0.0.1:20000")
 	if err != nil{
 		fmt.Printf("本地端口启动服务err:%v\n",err)
 		return
 	}
 	defer listener.Close()
 	// 2 等待别人来跟我建立连接
 	for{
		conn,err := listener.Accept()
	 	if err != nil{
	 		fmt.Printf("建立连接err：%v",err)
	 		return
	 	}
 		// 3 与客户端通信
    	go processConn(conn)	 	
 	}
}

func processConn(conn net.Conn){
	defer conn.Close()
	var temp [128]byte
	for{
		n,err := conn.Read(temp[:])
	 	if err != nil{
	 		fmt.Printf("通信err:%v",err)
	 	}
	 	fmt.Println(string(temp[:n]))
	}
 	
}