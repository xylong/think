package goroutine

import (
	"sync"
)

// Sum 求和计算
type Sum struct {
}

func NewSum() *Sum {
	return &Sum{}
}

// Sum 求和
func (s *Sum) Sum(min, max int) int {
	result := 0
	for i := min; i <= max; i++ {
		result = result + i
	}
	return result
}

// GoSum 按照group分批计算从min到max的和
func (s *Sum) GoSum(min, max, routineNum int) int {
	c := make(chan int, routineNum)
	wg := sync.WaitGroup{}

	result := 0
	n := max / routineNum

	wg.Add(routineNum)
	for i := 0; i < routineNum; i++ {
		go func(i int) {
			defer wg.Done()
			c <- s.Sum(n*i+1, n*i+n)
		}(i)
	}

	// 关闭channel防止range读取死锁
	go func() {
		defer close(c)
		wg.Wait()
	}()

	// 阻塞获取
	for v := range c {
		result = result + v
	}

	return result
}
