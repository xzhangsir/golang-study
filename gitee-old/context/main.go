package main

import (
	"fmt"
	"time"
	"sync"
	"context"
)

var wg sync.WaitGroup
func f(ctx context.Context){
LOOP:
	for{
		fmt.Println("haha")
		time.Sleep(time.Second)
		select{
		case  <- ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
	wg.Done()
}

func main(){
	ctx,cancel := context.WithCancel(context.Background())
    wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second*3)
	cancel()  // 通知子goroutine结束
	wg.Wait()
}