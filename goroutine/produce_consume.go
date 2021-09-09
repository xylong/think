package goroutine

import (
	"fmt"
	"time"
)

type PC struct {
	c chan int
}

func NewPC() *PC {
	return &PC{
		c: make(chan int),
	}
}

// 生产者
func (pc *PC) producer() {
	defer close(pc.c) // 不关闭的话消费者会一直执行，无法退出

	for i := 0; i < 5; i++ {
		pc.c <- i * 10
		time.Sleep(time.Millisecond * 500)
	}
}

// 消费者
// r 退出信号
func (pc *PC) consumer() (r chan struct{}) {
	r = make(chan struct{})

	go func() {
		defer func() {
			r <- struct{}{}
		}()

		for item := range pc.c {
			fmt.Println(item)
		}
	}()

	return r
}

func (pc *PC) Run() {
	go pc.producer()
	<-pc.consumer()
}
