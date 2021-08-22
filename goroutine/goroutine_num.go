package goroutine

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type goroutine struct {
}

func newGoroutine() *goroutine {
	return &goroutine{}
}

// job 模拟耗时任务
func (g *goroutine) job(index int) {
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("%d 执行完毕\n", index)
}

// SimpleLimit 简单限制协程数量
// num 任务数
// max 协程数
func (g *goroutine) SimpleLimit(num, max int) {
	wg := sync.WaitGroup{}
	pool := make(chan struct{}, max)

	// 100个任务，保证最多只有max个goroutine执行
	for i := 0; i < num; i++ {
		// 到达虽大长度阻塞
		pool <- struct{}{}

		wg.Add(1)
		go func(index int) {
			defer func() {
				wg.Done()
				<-pool
			}()
			g.job(index)
		}(i)
	}

	wg.Wait()
}

// Cycle 周期执行
func (g *goroutine) Cycle(max int) {
	pool := make(chan struct{}, max)

	go func() {
		for {
			for i := 0; i < max; i++ {
				pool <- struct{}{}
			}
			time.Sleep(time.Second * 3)
		}
	}()

	for {
		randNum := rand.Intn(100)
		<-pool
		go g.job(randNum)
	}
}
