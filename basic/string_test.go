package basic

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestString_Length(t *testing.T) {
	Convey("字符串长度", t, func() {
		So(From("").Length(), ShouldEqual, 0)
		So(FromInt(123).Length(), ShouldEqual, 3)
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
		So(From("我爱你 ok 123").Reverse(), ShouldEqual, "321 ko 你爱我")
		So(From("一2三4五6七8⑨").Reverse(), ShouldEqual, "⑨8七6五4三2一")
	})
}

func TestString_Reverse1(t *testing.T) {
	tests := []struct {
		name string
		s    String
		want string
	}{
		{"中文", "我爱你", "你爱我"},
		{"英文", "i love you", "uoy evol i"},
		{"数字", "123456789", "987654321"},
		{"特殊符号", "!@#$%^&*", "*&^%$#@!"},
		{"混用", "我!@#$%^&*you38", "83uoy*&^%$#@!我"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Reverse(); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
