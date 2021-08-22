package simple

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	Convey("两数之和", t, func() {
		nums, target, expectedValue := [][]int{{2, 7, 11, 15}, {3, 2, 4}, {3, 3}}, []int{9, 6, 6}, [][]int{{0, 1}, {1, 2}, {0, 1}}
		sum := 0

		for index, item := range nums {
			if reflect.DeepEqual(TwoSum(item, target[index]), expectedValue[index]) {
				sum++
			}
		}

		So(sum, ShouldEqual, len(expectedValue))
	})
}
