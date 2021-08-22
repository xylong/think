package question

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDefer_Param(t *testing.T) {
	Convey("defer定义了函数参数", t, func() {
		NewDefer().DefinedParam()
	})
}

func TestDefer_UndefinedParam(t *testing.T) {
	Convey("defer未定义函数参数", t, func() {
		NewDefer().UndefinedParam()
	})
}

func TestDefer_BreakThrough(t *testing.T) {
	Convey("突破defer函数参数机制", t, func() {
		NewDefer().BreakThrough()
	})
}

func TestDefer_Chain(t *testing.T) {
	Convey("defer链式调用", t, func() {
		d := NewDefer()
		defer d.Echo(1).Echo(2)
		d.Echo(3)
	})
}

func TestDefer_Loop(t *testing.T) {
	Convey("defer循环调用", t, func() {
		d := NewDefer()
		d.Loop()
	})
}
