package question

type Item struct {
	value    string
	priority int // 优先级
	index    int
}

type PriorityQueue []*Item

func (q *PriorityQueue) Len() int {
	return len(*q)
}

func (q *PriorityQueue) Less(i, j int) bool {
	queue := *q
	return queue[i].priority > queue[j].priority
	//return *q[i].priority > *q[j].priority
}

func (q *PriorityQueue) Swap(i, j int) {
	queue := *q
	queue[i], queue[j] = queue[j], queue[i]
	queue[i].index = i
	queue[j].index = j
}

func (q *PriorityQueue) Push(x any) {
	n := len(*q)
	item := x.(*Item)
	item.index = n
	*q = append(*q, item)
}

func (q *PriorityQueue) Pop() any {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // 避免内存泄露
	item.index = -1 // 安全标记
	*q = old[0 : n-1]
	return item
}
