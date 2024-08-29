package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"bufio"
)

func main(){
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 20000,
	})
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()
	// 发送数据
	sendData := bufio.NewReader(os.Stdin)
	for{
		input, _ := sendData.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if inputInfo == "exit"{
			break
		}
		_, err = socket.Write([]byte(inputInfo)) 
		if err != nil {
			fmt.Println("发送数据失败，err:", err)
			return
		}
		data := make([]byte, 4096)
		n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
		if err != nil {
			fmt.Println("接收数据失败，err:", err)
			return
		}
		fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
	}
}