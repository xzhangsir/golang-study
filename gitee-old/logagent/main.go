package main

import (
	"fmt"
	"time"
	"logagent/taillog"
	"logagent/kafka"
	"logagent/conf"
	"gopkg.in/ini.v1"
)
var cfg = new(conf.AppConf)
func run(){
	// 1 读取日志
	for {
		select {
		case line := <- taillog.ReadChan():
			// 2 发送到kafka
			kafka.SendMsgToKafka(cfg.KafkaConf.Topic,line.Text)
		default:
			time.Sleep(time.Second)
		}
	}	
}

// logagent入口程序
func main(){
	// 0 加载配置文件
    err := ini.MapTo(cfg,"./conf/config.ini")
    if err != nil{
    	fmt.Printf("load ini failed ,err:%v",err)
    	return
    }
	// 1   初始化kafka连接
  err = kafka.Init([]string{cfg.KafkaConf.Address})
  if err != nil{
  	fmt.Printf("init kafka failed,err:%v\n",err)
  	return
  }
  fmt.Println("kafka init success")
  // 2 打开日志文件准备收集日志
  err = taillog.Init(cfg.TaillogConf.Path)
  if err != nil{
  	fmt.Printf("init taillog failed,err:%v\n",err)
  	return
  }
  fmt.Println("taillog init success")
  run()
}