package decorator

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSum(t *testing.T) {
	Convey("累加求和", t, func() {
		So(Sum(1, 100000000), ShouldEqual, 5000000050000000)
	})
}

func TestConsume(t *testing.T) {
	Convey("累加耗时", t, func() {
		So(Consume(Sum)(1,100000000), ShouldEqual, 5000000050000000)
	})
}
