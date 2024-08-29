package main

import (
	"fmt"
	// "os"
	"flag"
)

// os.Args来获取命令行参数  返回切片类型
// go  run main.go 1  执行
/*func main(){
 	fmt.Printf("%#v\n",os.Args)
 	fmt.Println(os.Args[0])
 	fmt.Println(os.Args[1])
}*/

// go run main -name=xxx  执行
func main(){
						// flag名, 默认值, 帮助信息
	name := flag.String("name","zx","姓名")
	flag.Parse()  	//解析命令行参数
	fmt.Println(*name)

	fmt.Println(flag.Args())  ////返回命令行参数后的其他参数，以[]string类型
	fmt.Println(flag.NArg())  //返回命令行参数后的其他参数个数
	fmt.Println(flag.NFlag()) //返回使用的命令行参数个数
}