package question

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConcurrentWithoutLock(t *testing.T) {
	Convey("无锁的并发", t, func() {
		a, b := ConcurrentWithoutLock(10), ConcurrentWithoutLock(100000)
		fmt.Println(a, b)
		So(b, ShouldEqual, a)
	})
}

func TestConcurrentWithLock(t *testing.T) {
	Convey("有锁的并发", t, func() {
		a, b := ConcurrentWithLock(10), ConcurrentWithLock(100000)
		fmt.Println(a, b)
		So(b, ShouldEqual, a)
	})
}
