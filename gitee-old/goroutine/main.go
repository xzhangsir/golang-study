package main;

import (
	"fmt"
	"sync"
	// "strconv"
	// "time"
	// "math/rand"
	"sync/atomic"
)

var wg sync.WaitGroup

// goroutine实现并发
/*func hello(i int){
	defer wg.Done() //  这个goruntine 结束就 减1
	fmt.Println(i)
}
func main(){
	for i := 0 ; i < 10 ;i++{
		wg.Add(1)    //每次启动一个 goroutine  就加 1
		go hello(i)  //启动一个goroutine
	}
	wg.Wait() // 等待所有的goruntine结束
}*/


// 无缓存区的通道 单向通道
/*func recv(ch chan int){
   x := <- c
   fmt.Println("接受值成功",x)
}
func main(){
   var ch chan int
   ch = make(chan int)
   go recv(ch)   // 启用goroutine从通道接收值
   ch <- 10  //无缓冲的通道只有在有人接收值的时候才能发送值
   fmt.Println("发送成功",ch)
}
*/

// 有缓存区的通道
/*func main(){
	var ch chan int
	ch = make(chan int,1)
	ch <- 10  //接受值
	x := <- ch //发送值
	fmt.Println(x)
}
*/

// 从通道中取值

/*func main(){
	ch1 := make(chan int,10)
	ch2 := make(chan int,10)
	wg.Add(2)
   // 开启goroutine将0~10的数发送到ch1中
	go func(){
		for i := 0 ; i < 10 ;i++{
			ch1 <- i
		}
		wg.Done()
		close(ch1)  //通道关闭后 只能读取 不能写入
		}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func(){
		for {
			i , ok := <- ch1   // 通道关闭后再取值ok=false
			if !ok{
				break
			}
			ch2 <- i * i
		}
		wg.Done()
		close(ch2)
	}()
	wg.Wait()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2{  // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}*/

// 从通道中取值 
// <-chan int  只读单向通道   chan<- int  只写单向通道
/*func write(ch chan<- int){
	defer wg.Done()
 	for i:=0;i<10;i++{
 		ch <- i
 	}
 	close(ch)
}
func read(ch1 <-chan int,ch2 chan<- int){
	defer wg.Done()
	for i := range ch1{
		ch2 <- i*i
	}
	close(ch2)
}
func main(){
	ch1 := make(chan int,10)
	ch2 := make(chan int,10)
    wg.Add(2)
    go  write(ch1)
    go  read(ch1,ch2)
    for i := range ch2{
    	fmt.Println(i)
    }
    wg.Wait()
}*/


// worker pool   goroutine池
// 3个 goruntine  去执行5个任务
/*func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	// 输出结果
	for a := 1; a <= 5; a++ {
		<-results
	}
}*/



/*使用goroutine和channel实现一个计算int64随机数各位数和的程序。
 1 开启一个goroutine循环生成int64类型的随机数，发送到jobChan
 2 开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
 3 主goroutine从resultChan取出结果并打印到终端输出*/
/*type job struct{
	num int64
}
type result struct{
	job *job
	res int64
}
func randNum(ch chan<- *job){
	defer wg.Done()
	for{
	  x := rand.Int63()
	  newJob := &job{
	  	num:x,
	  }
      ch <- newJob
      time.Sleep(time.Second)
	}
}
func add(ch1 <-chan *job ,ch2 chan<- *result){
	defer wg.Done()
	for{
		job := <- ch1
		newResult := &result{
			job:job,
			res:randNumAdd(job.num),
		}	 
		ch2 <- newResult
	}
}
func randNumAdd(i int64)(sum int64){
	for i > 0 {
		a := i % 10
		i = i/10
		sum += a
	}
	return
}
func main(){
	jobChan := make(chan *job,100)
	resultChan := make(chan *result,100)
	wg.Add(24)
	go randNum(jobChan)
	for i := 0 ;i < 24 ;i++{
		go add(jobChan,resultChan)
	}
	for i := range resultChan{
		fmt.Printf("随机数%d,和值%d\n",i.job.num,i.res)
	}
	wg.Wait()

}*/

//互斥锁  能够保证同时只有一个goroutine可以访问共享资源
/*var lock sync.Mutex
var x = 0
func add(){
	for i := 0 ;i < 10000 ;i++{
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}
func main(){
   wg.Add(2)
   go add()
   go add()
   wg.Wait()
   fmt.Println(x)
}*/


// 读写互斥锁
// 当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
// 当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待

/*var rwlock sync.RWMutex

rwlock.RLock()   //加读锁
rwlock.RUnlock() //解读锁
rwlock.Lock()   // 加写锁
rwlock.Unlock()  // 解写锁*/


// sync.Once  某一项任务只执行一次
// var once sync.Once
// 例如关闭通道  once.Do(func(){close(ch1)})


// sync.Map
// go语言内置的map 大并发下可能会有问题
// 所有go语言 提供了一个开箱即用的并发安全版map–sync.Map


/*var m = sync.Map{}

func main(){
	for i:= 0 ; i < 20 ;i++{
		wg.Add(1)
		go  func(n int){
			key := strconv.Itoa(n)
			m.Store(key,n)  //存值
			val,_:=m.Load(key) //取值
			fmt.Printf("k=:%v,v:=%v\n", key, val)
			wg.Done()
		}(i)
	}
	wg.Wait()
}*/


// 原子操作  atomic包

/*var x int64 = 0
// var lock sync.Mutex
func add(){
	// lock.Lock()
	// x = x+1
	// lock.Unlock()
	atomic.AddInt64(&x,1)
    wg.Done()
}
func main(){
	wg.Add(1000)
   for i := 0 ;i < 1000 ;i++{
       go add()
   }
   wg.Wait()
   fmt.Println(x)
}*/