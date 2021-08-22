package goroutine

import (
	"fmt"
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
func (g *goroutine) SimpleLimit(max int) {
	wg := sync.WaitGroup{}
	pool := make(chan struct{}, max)

	for i := 0; i < 100; i++ {
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
