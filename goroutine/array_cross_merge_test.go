package goroutine

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var (
	fruits = []string{"apple", "banana", "cherry", "durian"}
	prices = []int{1, 2, 3, 4}
)

func TestRun(t *testing.T) {
	Convey("交叉合并数组", t, func() {
		a := NewArrayCrossMerge()
		ShouldEqual(a.Run(fruits, prices), []interface{}{"apple", 1, "banana", 2, "cherry", 3, "durian", 4})
	})
}

func TestArrayCrossMerge_Go(t *testing.T) {
	Convey("协程交叉合并数组", t, func() {
		a := NewArrayCrossMerge()
		ShouldEqual(a.Go(fruits, prices), []interface{}{"apple", 1, "banana", 2, "cherry", 3, "durian", 4})
	})
}
