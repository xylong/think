package goroutine

import (
	"fmt"
	"think/factory"
)

var (
	min   = 1
	max   = 100000000
	group = 10
)

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
	s.c <- s.sum1(min, max)
}

func (s *Sum) Run() {
	s.Start()
	result := s.sum1(min, max)
	s.End()

	fmt.Println(result)
}

func (s *Sum) Go() {
	s.Start()
	s.c = make(chan int, group)

	result := 0
	for i := 0; i < group; i++ {
		n := max / group
		s.sum2(n*i, n*i+n)
	}
	// 关闭channel，不然range会死锁
	close(s.c)

	for r := range s.c {
		result = result + r
	}

	s.End()

	fmt.Println(result)
}
