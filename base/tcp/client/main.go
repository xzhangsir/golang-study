package main;

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

func main(){
	// 与server端建立连接
	conn,err := net.Dial("tcp","127.0.0.1:20000")
	if err != nil{
		fmt.Printf("建立连接错误%v",err)
		return
	}
	// 发送数据
	inputReader := bufio.NewReader(os.Stdin)
	for{
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if inputInfo == "exit"{
			break
		}
		conn.Write([]byte(inputInfo))
	}
	// 关闭连接
	conn.Close()
}