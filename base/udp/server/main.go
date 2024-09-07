package main

import (
	"fmt"
	"net"
)

func main(){
  listen,err := net.ListenUDP("udp",&net.UDPAddr{
  	IP:   net.IPv4(127, 0, 0, 1),
  	Port:20000,
  })
  if err != nil{
  	fmt.Printf("err%v",err)
  	return
  }
  defer listen.	Close()
  for{
  	var data [1024]byte
  	n,addr,err := listen.ReadFromUDP(data[:])  // 接收数据
  	if err != nil{
  		fmt.Printf("err%v",err)
  		continue
  	}
  	fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
  	_,err = listen.WriteToUDP(data[:n],addr) // 发送数据
  	if err != nil {
		fmt.Println("write to udp failed, err:", err)
		continue
	}
  }
}