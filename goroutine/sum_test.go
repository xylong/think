package goroutine

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

var (
	// 最小值
	min = 1
	// 最大值
	max = 100000000
	// 协程数
	routineNum = 10
)

func TestSum_Sum(t *testing.T) {
	Convey("单线程求和", t, func() {
		sum := NewSum()
		start := time.Now()
		result := sum.Sum(min, max)
		end := time.Now()
		So(result, ShouldEqual, 5000000050000000)
		t.Log(end.Sub(start))
	})
}

func TestSum_Go(t *testing.T) {
	Convey("多协程求和", t, func() {
		sum := NewSum()
		start := time.Now()
		result := sum.GoSum(min, max, routineNum)
		end := time.Now()
		So(result, ShouldEqual, 5000000050000000)
		t.Log(end.Sub(start))
	})
}
