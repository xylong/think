package goroutine

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var (
	// 最小值
	min = 1
	// 最大值
	max = 100000000
	// 协程数
	routineNum = 10
	// 求和结果和
	result = 5000000500000000
)

func TestSum_Run(t *testing.T) {
	Convey("单线程求和", t, func() {
		sum := NewSum()
		ShouldEqual(sum.Run(min, max), result)
		t.Log(sum.TakeUpTime())
	})
}

func TestSum_Go(t *testing.T) {
	Convey("多协程求和", t, func() {
		sum := NewSum()
		ShouldEqual(sum.Go(min, max, routineNum), result)
		t.Log(sum.TakeUpTime())
	})
}
