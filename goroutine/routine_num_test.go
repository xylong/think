package goroutine

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGoroutine_SimpleLimit(t *testing.T) {
	Convey("限制协程数量的基本方法", t, func() {
		limit := 10
		newGoroutine().SimpleLimit(limit)
	})
}
