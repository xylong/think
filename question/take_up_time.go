package question

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// MyTime 统计多个协程耗时时间
type MyTime struct {
	sync.WaitGroup
	sync.Mutex
	t int64 // 累加时间
}

// Done 任务完成
// 累加执行任务耗时
func (mt *MyTime) Done(start time.Time) {
	defer mt.WaitGroup.Done()
	defer func() {
		mt.Lock()
		mt.t += time.Since(start).Milliseconds() / 1000
		mt.Unlock()
	}()
}

func NewMyTime() *MyTime {
	return &MyTime{}
}

// job 模拟任务耗时
func (mt *MyTime) job() {
	t := time.Second * time.Duration(rand.Intn(4))
	time.Sleep(t)
	fmt.Printf("运行了%s秒\n", t.String())
}

// TakeUp 耗时
func (mt *MyTime) TakeUp() {
	start := time.Now()

	for i := 0; i < 5; i++ {
		mt.Add(1)
		go func() {
			defer mt.Done(time.Now())
			mt.job()
		}()
	}

	mt.Wait()
	fmt.Println("主线程执行：", time.Since(start).Milliseconds()/1000)
	fmt.Println("协程实际总共运行：", mt.t)
}
