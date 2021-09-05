package basic

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestString_Length(t *testing.T) {
	Convey("字符串长度", t, func() {
		ShouldEqual(From("").Length(), 0)
		ShouldEqual(FromInt(123).Length(), 3)
	})
}

func TestString_Each(t *testing.T) {
	Convey("遍历字符串", t, func() {
		From("你我他 golang 123").Each(func(item string) {
			fmt.Println(item)
		})
	})
}

func TestString_Reverse(t *testing.T) {
	Convey("反转字符串", t, func() {
		ShouldEqual(From("你爱我 ok 123").Reverse(), "321 ok 我爱你")
	})
}
