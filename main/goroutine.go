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


//----channel类型---

// var 变量名称 chan  元素类型
// var ch2 chan bool  // 声明一个传递布尔型的通道
// var ch3 chan []int // 声明一个传递int切片的通道

// 初始化channel  
// make(chan 元素类型,[缓存大小])

// 通道有 发送 接受 关闭三种操作
// ch := make(chan int)
// // 发送 
// ch <- 10 //把10发送到ch中
// // 接受 
// x := <- ch  // 从ch中接收值并赋值给变量x
// <- ch   //从ch中接受值，忽略结果
// // 关闭
// close(ch)

// 无缓冲通道（同步通道）
// var wg sync.WaitGroup
// func recv(c chan int){
// 	defer wg.Done()
// 	x := <-c
// 	fmt.Println("接收成功", x)
// }
// func main(){
// 	ch := make(chan int)
// 	wg.Add(1)
// 	go recv(ch) //创建一个 goroutine 从通道接收值
// 	ch <- 10
// 	fmt.Println("发送成功")
// 	wg.Wait()
// }


// 有缓冲的通道
// func main(){
// 	ch := make(chan int,1)
// 	ch <- 10
// 	fmt.Println("发送成功")
// }

// 判断通道是否关闭
// value,ok := <- ch
// func f2(c chan int){
// 	for {
// 		v,ok := <- c
// 		if !ok{
// 			fmt.Println("通道关闭了")
// 			break;
// 		}
// 		fmt.Printf("v:%#v ok:%#v\n", v, ok)
// 	}
// }

// for range接收值
// 通道内的所有值被接收完毕后会自动退出循环
// func f2(c chan int){
// 	for v := range c{
// 		fmt.Println(v)
// 	}
// }

// func main(){
// 	ch := make(chan int,2)
// 	ch <- 10
// 	ch <- 20
// 	close(ch)
// 	f2(ch)
// }

// 单向通道
// <- chan int // 只接收通道，只能接收不能发送
// chan <- int // 只发送通道，只能发送不能接收

// Producer 返回一个接收通道
// 并持续将符合条件的数据发送至返回的通道中
// 数据发送完成后会将返回的通道关闭
func Producer() <-chan int{
	ch := make(chan int,2)
	go func(){
		for i := 0 ; i < 10 ; i++{
			if i%2 == 1{
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}
// Consumer 从通道中接收数据进行计算 参数为接受通道
func Consumber(ch <-chan int) int{
	sum := 0
	for v := range ch{
		sum += v
	}
	return sum
}

func main(){
	ch := Producer()
	res := Consumber(ch)
	fmt.Println(res) //25
}

