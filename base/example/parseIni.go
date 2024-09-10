package example

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}
type Config struct {
	MysqlConfig `ini:"mysql"`
}

func loadIni(data interface{}) (err error) {
	// 参数校验
	// 传进来的 data 必须是指针类型
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = fmt.Errorf("data 不是指针类型")
		return
	}
	// 并且data必须是结构体类型的指针
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data 不是结构体类型的指针")
		return
	}
	// 读取配置文件的所有内容 得到字节类型的数据
	b, err := os.ReadFile("./temp/mysql.ini")
	if err != nil {
		fmt.Println("读取错误:", err)
		return
	}
	// 将字节类型转化为string类型 一行一行的切割
	lineSlice := strings.Split(string(b), "\r\n")
	var structName string
	// 一行一行的读取数据
	for idx, line := range lineSlice {
		// 去掉字符串首位的空格
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		// 如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 如果是[开头，并且]结尾的就是节
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("语法错误 %d", idx+1)
				return
			}
			// 并且 [] 中间要有内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("语法错误 %d", idx+1)
				return
			}
			// 根据节中的内容去data里面根据反射找到对应的结构体
			//  通过for循环遍历结构体的所有字段信息
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if field.Tag.Get("ini") == sectionName {
					structName = field.Name
					break
				}
			}
		} else {
			if len(structName) == 0 {
				return errors.New("没有对应的结构体")
			} else {
				//  = 来分割的键值对
				if !strings.Contains(line, "=") {
					err = fmt.Errorf("语法错误%d", idx+1)
					return
				}
				index := strings.Index(line, "=")
				key := strings.TrimSpace(line[:index])
				value := strings.TrimSpace(line[index+1:])
				// 根据structName去data中吧对应的嵌套结构体取出来
				v := reflect.ValueOf(data)
				// 拿到嵌套结构体的值信息
				sVal := v.Elem().FieldByName(structName)
				// 拿到嵌套结构体的类型
				sType := sVal.Type()
				if sType.Kind() != reflect.Struct {
					err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
					return
				}
				var fieldName string
				var fileType reflect.StructField
				// 遍历结构体的每一个字段 判断tag是不是等于key
				for i := 0; i < sVal.NumField(); i++ {
					field := sType.Field(i)
					fileType = field
					if field.Tag.Get("ini") == key {
						fieldName = field.Name
						break
					}
				}
				// 根据fieldName 取出这个字段
				fileObj := sVal.FieldByName(fieldName)
				// fmt.Println(fieldName,fileObj.Type())
				switch fileType.Type.Kind() {
				case reflect.String:
					fileObj.SetString(value)
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					var valueInt int64
					valueInt, err = strconv.ParseInt(value, 10, 64)
					if err != nil {
						return
					}
					fileObj.SetInt(valueInt)
				}
			}
		}
	}
	return
}

// 解析ini文件
func ParseIni() {
	var fc Config
	err := loadIni(&fc)
	if err != nil {
		fmt.Printf("load ini err %s", err)
		return
	}
	fmt.Println(fc.Address, fc.Port, fc.Username, fc.Password)
}
