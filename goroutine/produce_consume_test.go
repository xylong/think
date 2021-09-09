package goroutine

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPC_Run(t *testing.T) {
	Convey("生产消费模式", t, func() {
		NewPC().Run()
	})
}
