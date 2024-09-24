package base

func GoroutineFunc() {
	// singleRun()
	// multRun()
	// 通道
	// channelFunc()
	// context
	// contextFunc()
}

/* //----启动单个goroutine----
var wg sync.WaitGroup

func hello() {
	defer wg.Done()
	fmt.Println("hello")
}
func singleRun() {
	wg.Add(1)
	go hello()
	fmt.Println("你好")
	wg.Wait()
}
*/
// ----启动多个goroutine---

/* var wg sync.WaitGroup

func hello(v int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("hello", v)
}
func multRun() {
	for v := range 10 {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(v)
	}
	fmt.Println("你好")
	wg.Wait() // 等待所有登记的goroutine都结束
}
*/
/*
// ----channel类型-通道--

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
// func recv(ch chan int) {
// 	x := <-ch
// 	fmt.Println("接受成功", x)
// }
// func channelFunc() {
// 	ch := make(chan int)
// 	go recv(ch)
// 	ch <- 10
// 	fmt.Println("发送成功")
// }

// 有缓冲的通道
// func channelFunc(){
// 	ch := make(chan int,1)
// 	ch <- 10
// 	fmt.Println("发送成功")
// }

// 关闭通道
// func f2(ch chan int) {
// 	// for {
// 	// 	v, ok := <-ch
// 	// 	if !ok {
// 	// 		fmt.Println("通道已关闭")
// 	// 		break
// 	// 	}
// 	// 	fmt.Println(v)
// 	// }
//
// 	for v := range ch {
// 		fmt.Println(v)
// 	}
// }

// func channelFunc() {
// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2
// 	close(ch)
// 	f2(ch)
// }

// <- chan int // 只接收通道，只能接收不能发送
// chan <- int // 只发送通道，只能发送不能接收
// func producer() <-chan int {
// 	ch := make(chan int, 2)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			if i%2 == 1 {
// 				ch <- i
// 			}
// 		}
// 		close(ch)
// 	}()
// 	return ch
// }
// func consumer(ch <-chan int) int {
// 	var sum = 0
// 	for v := range ch {
// 		sum += v
// 	}
// 	return sum
// }

// func channelFunc() {
// 	ch := producer()
// 	sum := consumer(ch)
// 	fmt.Println(sum) //25
// }

// select多路复用
// func channelFunc() {
// 	ch := make(chan int, 1)

// 	for i := 0; i < 10; i++ {
// 		select {
// 		case ch <- i:
// 		case x := <-ch:
// 			fmt.Println(x)
// 		}
// 	}
// }

// 互斥锁
// var (
// 	x  int
// 	wg sync.WaitGroup //等待组
// 	m  sync.Mutex     //互斥锁
// )

// func add() {
// 	for i := 0; i < 5000; i++ {
// 		m.Lock() //修改X前加锁
// 		x += 1
// 		m.Unlock() //修改X后解锁
// 	}
// 	wg.Done()
// }

// func channelFunc() {
// 	wg.Add(2)
// 	go add()
// 	go add()
// 	wg.Wait()
// 	fmt.Println(x)
// }

//  读写锁  sync.RWMutex
// var rw sync.RWMutex
// rw.Lock() //	获取写锁
// rw.Unlock() //释放写锁
// rw.RLock() //获取读锁
// rw.RUnlock() //释放读锁
//当一个 goroutine 获取到读锁之后，其他的 goroutine 如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
// 而当一个 goroutine 获取写锁之后，其他的 goroutine 无论是获取读锁还是写锁都会等待

// 只执行一次
// var one sync.Once
// one.Do(xxx)

// 并发安全的 sync.Map
// func channelFunc() {
// 	m := sync.Map{}
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func(n int) {
// 			defer wg.Done()
// 			key := strconv.Itoa(n) // 整数转字符串
// 			m.Store(key, n)
// 			v, _ := m.Load(key)
// 			fmt.Printf("k=:%v,v:=%v\n", key, v)
// 		}(i)
// 	}
// 	wg.Wait()
// }
*/

/*----context-----*/

/*
var wg sync.WaitGroup

	func doTash(n int) {
		time.Sleep(time.Duration(n))
		fmt.Printf("Task %d Done\n", n)
		wg.Done()
	}

	func contextFunc() {
		for i := 0; i < 3; i++ {
			go doTash(i + 1)
			wg.Add(1)
		}
		// wg.Wait() 会等待所有的子协程任务全部完成，所有子协程结束后，才会执行 wg.Wait() 后面的代码
		// 如何主动通知子协程退出？ 下一个例子
		wg.Wait()
		fmt.Println("主结束")
	}
*/

/*
	// select+chan

// 子协程使用 for 循环定时轮询，如果 stop 信道有值，则退出，否则继续轮询
var stop chan bool

	func reqTask(name string) {
		for {
			select {
			case <-stop:
				fmt.Println("stop", name)
				return
			default:
				fmt.Println(name, "send request")
				time.Sleep(1 * time.Second)
			}
		}
	}

	func contextFunc() {
		stop = make(chan bool)
		go reqTask("work1")
		time.Sleep(3 * time.Second)
		stop <- true
		time.Sleep(3 * time.Second)
		fmt.Println("主结束")
	}
*/
/* func reqTask(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			op := ctx.Value("options")
			if op != nil {
				fmt.Println(name, "send request", op)
			} else {
				fmt.Println(name, "send request")
			}

			time.Sleep(1 * time.Second)
		}
	}
}

func contextFunc() {
	ctx, cancal := context.WithCancel(context.Background())
	go reqTask(ctx, "work1")
	// 往子协程中传参
	ctxv := context.WithValue(ctx, "options", "我是值")
	go reqTask(ctxv, "work2")
	time.Sleep(3 * time.Second)
	cancal()
	time.Sleep(3 * time.Second)
	fmt.Println("主结束")
} */
