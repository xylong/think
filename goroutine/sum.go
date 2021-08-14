package goroutine

import (
	"fmt"
	"sync"
	"think/factory"
)

var (
	min   = 1
	max   = 100000000
	group = 2
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
	r := s.sum1(min, max)
	s.c <- r
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

	go func() {
		defer close(s.c)
		wg.Wait()
	}()

	for v := range s.c {
		result = result + v
	}

	s.End()

	fmt.Println(result)
}
