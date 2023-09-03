package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup等待执行
func wait() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("协程启动", time.Now().Format("15:04:05"))
		time.Sleep(time.Second * 2) // 业务执行
	}()

	fmt.Println("正常执行", time.Now().Format("15:04:05"))
	wg.Wait()
	fmt.Println("执行结束", time.Now().Format("15:04:05"))
}

// 通过channel等待
func waitByChannel() {
	ch := make(chan struct{})
	go func() {
		fmt.Println("协程启动", time.Now().Format("15:04:05"))
		time.Sleep(time.Second * 2)
		ch <- struct{}{}
	}()

	fmt.Println("正常执行", time.Now().Format("15:04:05"))
	<-ch
	fmt.Println("执行结束", time.Now().Format("15:04:05"))
}

// Pipeline 管道模式，即一个Channel的输出作为下一个Channel的输入
func Pipeline() {
	a := make(chan int)
	b := make(chan int)

	go func() {
		defer close(a)

		for i := 0; i < 10; i++ {
			a <- i
		}

	}()

	go func() {
		defer close(b)

		for i := range a {
			b <- i * i
		}
	}()

	for i := range b {
		fmt.Println(i)
	}
}

func main() {
	//wait()
	//waitByChannel()
	Pipeline()
}
