package simple

// 派去
type sort struct {
}

func newSort() *sort {
	return &sort{}
}

// QuickSort 快速排序
func (s *sort) QuickSort(arr []int) []int {
	// 处理空数组或一个元素的
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0] // 基准值
	var left, right []int

	// 比基准值小放左边，比基准值大放右边
	for _, item := range arr[1:] {
		if item <= pivot {
			left = append(left, item)
		} else {
			right = append(right, item)
		}
	}

	// QuickSort(左边) + pivot + QuickSort(右边)
	return append(s.QuickSort(left), append([]int{pivot}, s.QuickSort(right)...)...)
}
