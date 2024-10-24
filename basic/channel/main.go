package main

import (
	"fmt"
	"time"
)

// 读没有值的channel,如果是在主线程中则会死锁，如果是协程中则是阻塞
func readEmptyChannel() {
	q := make(chan int, 2)
	<-q
}

// 读关闭的channel，没有值也不会报错，只会读出0值和false
func readCloseChannel() {
	q := make(chan int, 2)
	close(q)
	v, ok := <-q
	fmt.Println(v, ok)
}

// 对关闭的channel写入操作会panic
func writeCloseChannel() {
	q := make(chan int, 2)
	close(q)
	q <- 1
}

func main() {
	readEmptyChannel()
	go readEmptyChannel()
	time.Sleep(time.Second * 2)
	fmt.Println("done")
	//readCloseChannel()
	//writeCloseChannel()
}
