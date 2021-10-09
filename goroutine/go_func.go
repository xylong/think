package goroutine

import "sync"

/*
 * 封装和简化协程+channel传值
 */

// TaskFn 任务执行的函数
type TaskFn func() interface{}

// Task 协程任务
type Task struct {
	data  chan interface{}
	wg    *sync.WaitGroup
	tasks []TaskFn
}

// NewTask 创建任务
func NewTask() *Task {
	return &Task{
		data: make(chan interface{}, 10),
		wg:   &sync.WaitGroup{},
	}
}

// Add 添加任务
func (t *Task) Add(f TaskFn) {
	t.tasks = append(t.tasks, f)
}

// 执行所有任务
func (t *Task) do() {
	for _, f := range t.tasks {
		t.wg.Add(1)
		go func() {
			defer t.wg.Done()
			t.data <- f()
		}()
	}
}

// Range 遍历任务结果
// f 回调函数，将结果通过回调函数返回
func (t *Task) Range(f func(i interface{})) {
	t.do()

	go func() {
		defer close(t.data)
		t.wg.Wait()
	}()

	for item := range t.data {
		f(item)
	}
}
