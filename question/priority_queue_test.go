package question

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPriorityQueue_Len(t *testing.T) {
	items := map[string]int{
		"任务1": 3,
		"任务2": 2,
		"任务3": 5,
	}

	queue := make(PriorityQueue, 0)
	for s, i := range items {
		queue.Push(&Item{
			value:    s,
			priority: i,
			index:    i,
		})
	}

	fmt.Println(queue.Len())
}

func TestPriorityQueue(t *testing.T) {
	items := map[string]int{
		"任务1": 3,
		"任务2": 2,
		"任务3": 5,
		"任务4": 1,
	}

	queue := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		queue[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&queue)

	item := &Item{
		value:    "任务5",
		priority: 4,
	}
	heap.Push(&queue, item)

	for queue.Len() > 0 {
		item := heap.Pop(&queue).(*Item)
		fmt.Printf("处理: %s\n", item.value)
	}
}
