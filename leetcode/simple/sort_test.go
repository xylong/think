package simple

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

var (
	arr []int = []int{12, 87, 1, 66, 30, 126, 328, 12, 653, 67, 98, 3, 256}
	brr []int = []int{1, 3, 12, 12, 30, 66, 67, 87, 98, 126, 256, 328, 653}
)

func TestSort_QuickSort(t *testing.T) {
	Convey("快速排序", t, func() {
		So(reflect.DeepEqual(newSort().QuickSort(arr), brr), ShouldBeTrue)
	})
}
