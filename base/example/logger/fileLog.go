package myLogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type fileLog struct {
	level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

func NewFileLog(levelStr, filePath, fileName string, maxFileSize int64) *fileLog {
	level, err := parseStringLogLevel(levelStr)
	if err != nil {
		panic("err")
	}
	f1 := &fileLog{
		level:       level,
		filePath:    filePath,
		fileName:    fileName,
		maxFileSize: maxFileSize,
	}
	err = f1.initFile()
	if err != nil {
		panic(err)
	}
	return f1
}
func (f *fileLog) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("写入文件错误%s", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Print("写入文件错误%s", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}
func (f *fileLog) log(lv LogLevel, format string, a ...interface{}) {
	if f.level <= lv {
		msg := fmt.Sprintf(format, a...)
		funcName, fileName, lineNo := getInfo(3)
		now := time.Now()
		lev := parseLogLevelString(lv)
		// f.fileObj.WriteString("直接写入字符串数据")
		newFile, errs := f.splitLogFile(f.fileObj)
		if errs != nil {
			return
		}
		f.fileObj = newFile
		fmt.Fprintf(f.fileObj, "[%s] [%s] 文件名[%s] 方法名[%s] 行数[%v] %s\n", now.Format("2006-01-02 15:04:05"), lev, fileName, funcName, lineNo, msg)
		if lv >= err {
			newErrFileObj, err := f.splitLogFile(f.errFileObj)
			if err != nil {
				return
			}
			f.errFileObj = newErrFileObj
			fmt.Fprintf(f.errFileObj, "[%s] [%s] 文件名[%s] 方法名[%s] 行数[%v] %s\n", now.Format("2006-01-02 15:04:05"), lev, fileName, funcName, lineNo, msg)
		}
	}
}

// 切割文件
func (f *fileLog) splitLogFile(file *os.File) (*os.File, error) {

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("获取文件信息失败啦%s", err)
		return nil, err
	}
	// 如果当前的文件大于 直接的文件大小 就进行切割
	if fileInfo.Size() >= f.maxFileSize {
		nowStr := time.Now().Format("20060102150405000")       //获取当前的时间
		logName := path.Join(f.filePath, fileInfo.Name())      //获取当前文件的路径
		newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr) //起一个新的名字
		file.Close()                                           //关闭当前的日志文件
		os.Rename(logName, newLogName)                         //对当前的文件重命名
		// 打开一个新的日志文件
		fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Print("打开新的日志文件错误")
			return nil, err
		}
		return fileObj, nil
	} else {
		return file, nil
	}

}

func (f *fileLog) Debug(format string, a ...interface{}) {
	f.log(debug, format, a...)
}
func (f *fileLog) Warning(format string, a ...interface{}) {
	f.log(warning, format, a...)
}
func (f *fileLog) Err(format string, a ...interface{}) {
	f.log(err, format, a...)
}
func (f *fileLog) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
