package main

import (
	"fmt"
	"strconv"
	"sync"
)

//----启动单个goroutine----

// 声明全局等待组变量
// var wg sync.WaitGroup

// func hello(){
// 	defer wg.Done() // 告知当前goroutine完成
// 	fmt.Println("hello")
// }

// func Goroutine(){

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
// func Goroutine(){
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
// func Goroutine(){
// 	ch := make(chan int)
// 	wg.Add(1)
// 	go recv(ch) //创建一个 goroutine 从通道接收值
// 	ch <- 10
// 	fmt.Println("发送成功")
// 	wg.Wait()
// }

// 有缓冲的通道
// func Goroutine(){
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

// func Goroutine(){
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
// func Producer() <-chan int{
// 	ch := make(chan int,2)
// 	go func(){
// 		for i := 0 ; i < 10 ; i++{
// 			if i%2 == 1{
// 				ch <- i
// 			}
// 		}
// 		close(ch)
// 	}()
// 	return ch
// }
// // Consumer 从通道中接收数据进行计算 参数为接受通道
// func Consumber(ch <-chan int) int{
// 	sum := 0
// 	for v := range ch{
// 		sum += v
// 	}
// 	return sum
// }

// func Goroutine(){
// 	ch := Producer()
// 	res := Consumber(ch)
// 	fmt.Println(res) //25
// }

// select多路复用
// func Goroutine(){
// 	ch := make(chan int,1)
// 	for i := 1 ; i <= 10 ; i++{
// 		select{
// 		case x := <- ch:
// 			fmt.Println(x) // 1 3 5 7 9
// 		case ch <- i:
// 		}
// 	}
// // 第一次循环时 i = 1，select 语句中包含两个 case 分支，此时由于通道中没有值可以接收，所以x := <-ch 这个 case 分支不满足，而ch <- i这个分支可以执行，会把1发送到通道中，结束本次 for 循环；
// // 第二次 for 循环时，i = 2，由于通道缓冲区已满，所以ch <- i这个分支不满足，而x := <-ch这个分支可以执行，从通道接收值1并赋值给变量 x ，所以会在终端打印出 1；
// // 后续的 for 循环以此类推会依次打印出3、5、7、9。
// }

// 互斥锁
// 如果不加锁
// 开启了两个 goroutine 分别执行 add 函数，
// 这两个 goroutine 在访问和修改全局的x变量时就会存在数据竞争，
// 某个 goroutine 中对全局变量x的修改可能会覆盖掉另一个 goroutine 中的操作，
// 所以导致最后的结果与预期不符。

// var (
// 	x int
// 	wg sync.WaitGroup //等待组
// 	m sync.Mutex //互斥锁
// )
// func add(){
// 	for i:=0 ; i < 5000 ; i++{
// 		m.Lock()  //修改X前加锁
// 		x = x + 1
// 		m.Unlock() //修改X后解锁
// 	}
// 	wg.Done()
// }
// func Goroutine(){
// 	wg.Add(2)
// 	go add()
// 	go add()
// 	wg.Wait()
// 	fmt.Println(x) //10000
// }

// sync.Map
// 不能在多个 goroutine 中并发对内置的 map 进行读写操作，
// 否则会存在数据竞争问题。

// var m = make(map[string]int)
var m = sync.Map{}

// func get(key string) int{
// 	return m[key]
// }
// func set(key string,v int){
// 	m[key] = v
// }

func Goroutine() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			// 整数转字符串
			// set(key, n)
			// fmt.Println(key,get(key))
			m.Store(key, i)         // 存储key-value
			value, _ := m.Load(key) // 根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
