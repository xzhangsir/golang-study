package main

import (
	"fmt"
	// "sync"
)

//----启动单个goroutine----

// 声明全局等待组变量
// var wg sync.WaitGroup

// func hello(){
// 	defer wg.Done() // 告知当前goroutine完成
// 	fmt.Println("hello")
// }

// func main(){

// 	wg.Add(1) // 登记1个goroutine
// 	go hello() // 启动另外一个goroutine去执行hello函数
// 	fmt.Println("我是main")
// 	wg.Wait() // 阻塞等待登记的goroutine完成

// }


// ----启动多个goroutine---
// var wg sync.WaitGroup
// func hello(i int){
// 	defer wg.Done() // goroutine结束就登记-1
// 	fmt.Println(i)
// }
// func main(){
// 	for i := 0 ; i < 10 ; i++{
// 		wg.Add(1) // 启动一个goroutine就登记+1
// 		go hello(i)
// 	}
// 	wg.Wait() // 等待所有登记的goroutine都结束
	
// }
