package question

import "sync"

// myLock 自定义锁
type myLock struct {
	ch chan struct{}
}

func NewMyLock() *myLock {
	return &myLock{
		ch: make(chan struct{}, 1), // ?缓冲区设成1防止阻塞
	}
}

func (ml *myLock) lock() {
	ml.ch <- struct{}{}
}

func (ml *myLock) unlock() {
	select {
	case <-ml.ch:
	// !防止多次解锁
	default:
		panic("unlock error")
	}
}

type stock struct {
	num  int
	lock *myLock
}

func (s *stock) incr() {
	s.num++
}

func (s *stock) decr() {
	s.num--
}

func (s *stock) increment() {
	defer s.lock.unlock()
	s.lock.lock()
	s.num++
}

func (s *stock) decrement() {
	defer s.lock.unlock()
	s.lock.lock()
	s.num--
}

// ConcurrentWithoutLock 没有锁的并发
// *连续加减相同的次数，库存变了
func ConcurrentWithoutLock(count int) int {
	s := &stock{}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < count; i++ {
			s.incr()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < count; i++ {
			s.decr()
		}
	}()

	wg.Wait()
	return s.num
}

// ConcurrentWithLock 有锁的并发
// *连续加减相同的次数，库存不变
func ConcurrentWithLock(count int) int {
	s := &stock{
		lock: NewMyLock(),
	}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < count; i++ {
			s.increment()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < count; i++ {
			s.decrement()
		}
	}()

	wg.Wait()
	return s.num
}
