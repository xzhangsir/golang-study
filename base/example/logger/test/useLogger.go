package main

import (
	myLogger "golang-study/base/example/logger"
	"time"
)

var log myLogger.Logger

func main() {
	log = myLogger.NewLogger("warning")
	// log = myLogger.NewFileLog("warning", "./", "zx.log", 2*1024)
	for {
		id := 12
		name := "zx"
		log.Debug("调试%d,%s", id, name)
		log.Warning("警告%d,%s", id, name)
		log.Err("错误%d,%s", id, name)
		// 每隔3秒执行一次
		time.Sleep(time.Second * 3)
	}

}
