package goroutine

import (
	"sync"
)

// ArrayCrossMerge 交叉合并数组
type ArrayCrossMerge struct {
}

func NewArrayCrossMerge() *ArrayCrossMerge {
	return &ArrayCrossMerge{}
}

func (acm *ArrayCrossMerge) Run(arr []string, brr []int) []interface{} {
	crr := make([]interface{}, 0)
	for index, value := range arr {
		crr = append(crr, value)
		crr = append(crr, brr[index])
	}

	return crr
}

func (acm *ArrayCrossMerge) Go(arr []string, brr []int) []interface{} {
	wg := sync.WaitGroup{}
	c1, c2 := make(chan struct{}), make(chan struct{})
	result := make([]interface{}, 0)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, value := range arr {
			<-c1
			result = append(result, value)
			c2 <- struct{}{}
		}
	}()

	wg.Add(1)
	go func() {
		for index, value := range brr {
			<-c2
			result = append(result, value)
			if index == len(brr)-1 {
				wg.Done()
			}
			c1 <- struct{}{}
		}
	}()
	c1 <- struct{}{}
	wg.Wait()
	return result
}
