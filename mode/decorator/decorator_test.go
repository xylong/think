package decorator

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDemoConsume(t *testing.T) {
	Convey("装饰器模式", t, func() {
		So(DemoConsume(demo)(10), ShouldEqual, 100)
	})
}
