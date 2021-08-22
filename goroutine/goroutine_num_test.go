package goroutine

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGoroutine_SimpleLimit(t *testing.T) {
	Convey("限制协程数量的基本方法", t, func() {
		num, limit := 100, 10
		newGoroutine().SimpleLimit(num, limit)
	})
}

func TestGoroutine_Cycle(t *testing.T) {
	Convey("控制协程数+周期性执行任务", t, func() {
		newGoroutine().Cycle(5)
	})
}
