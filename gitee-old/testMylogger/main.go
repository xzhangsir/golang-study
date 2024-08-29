package main;

import (
	"time"
	"mylogger"
)
var log mylogger.Loggerr
func main(){
   // log := mylogger.NewLogger("warning")
   log = mylogger.NewFileLog("warning","./","zx.log",2*1024)
   for{
   		id := 12
   		name := "zx"
   		log.Debug("调试%d,%s",id,name)
   		log.Warning("警告%d,%s",id,name)
   		log.Err("错误%d,%s",id,name)
   		// 每隔3秒执行一次
   		time.Sleep(time.Second * 0)
   }

}