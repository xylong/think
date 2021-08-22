package question

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMyTime_TakeUp(t *testing.T) {
	Convey("协程耗时", t, func() {
		NewMyTime().TakeUp()
	})
}
