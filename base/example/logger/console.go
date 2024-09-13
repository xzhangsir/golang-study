package myLogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

// 往终端写日志相关内容
type LogLevel uint16

const (
	undefined LogLevel = iota
	debug
	warning
	err
)

// Logger 日志结构体
type Logger struct {
	level LogLevel
}

// 将string类型转为LogLevel类型
func parseStringLogLevel(s string) (LogLevel, error) {
	switch s {
	case "debug":
		return debug, nil
	case "warning":
		return warning, nil
	case "err":
		return err, nil
	default:
		return undefined, errors.New("没有这个级别")
	}
}
func parseLogLevelString(l LogLevel) string {
	switch l {
	case debug:
		return "debug"
	case warning:
		return "warning"
	case err:
		return "err"
	default:
		return "undefined"
	}
}

// 日志结构体 构造函数
func NewLogger(levelStr string) Logger {
	level, err := parseStringLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		level: level,
	}
}

// 获取谁调用了我
func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Print("runtime.Caller() err")
		return
	}
	funcName = runtime.FuncForPC(pc).Name() //获取方法名
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return
}

func (l Logger) log(lv LogLevel, format string, a ...interface{}) {
	if l.level < lv {
		msg := fmt.Sprintf(format, a...)
		funcName, fileName, lineNo := getInfo(3)
		now := time.Now()
		lev := parseLogLevelString(lv)
		fmt.Printf("[%s] [%s] 文件名[%s] 方法名[%s] 行数[%v] %s\n", now.Format("2006-01-02 15:04:05"), lev, fileName, funcName, lineNo, msg)
	}
}

func (l Logger) Debug(format string, a ...interface{}) {
	l.log(debug, format, a...)
}
func (l Logger) Warning(format string, a ...interface{}) {
	l.log(warning, format, a...)
}
func (l Logger) Err(format string, a ...interface{}) {
	l.log(err, format, a...)
}
