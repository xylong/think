package goroutine

import (
	"fmt"
	"sync"
	"think/factory"
)

var (
	min   = 1
	max   = 100000000
	group = 10
)

// Sum 求和计算
type Sum struct {
	*factory.Demo
	c chan int
}

func NewSum() *Sum {
	return &Sum{Demo: &factory.Demo{}}
}

// sum1 求和
func (s *Sum) sum1(min, max int) int {
	result := 0
	for i := min; i <= max; i++ {
		result = result + i
	}
	return result
}

// sum2	求和
func (s *Sum) sum2(min, max int) {
	r := s.sum1(min, max)
	s.c <- r
}

// Run 单线程求和
func (s *Sum) Run() {
	s.Start()
	result := s.sum1(min, max)
	s.End()

	fmt.Println(result)
}

// Go 按照group分批计算从min到max的和
func (s *Sum) Go() {
	s.Start()
	s.c = make(chan int, group)
	wg := sync.WaitGroup{}

	result := 0
	wg.Add(group)
	for i := 0; i < group; i++ {
		go func(i int) {
			defer wg.Done()
			n := max / group
			s.sum2(n*i, n*i+n)
		}(i)
	}

	// 关闭channel防止range读取死锁
	go func() {
		defer close(s.c)
		wg.Wait()
	}()

	// 阻塞获取
	for v := range s.c {
		result = result + v
	}

	s.End()
	fmt.Println(result)
}
