package basic

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLuckyUser_Run(t *testing.T) {
	Convey("抽取幸运观众", t, func() {
		num := 9
		luckyDog := NewLuckyUser()

		luckyDog.setRange(1, 5).Run(num)
		t.Log(luckyDog.users)
		So(len(luckyDog.users), ShouldBeLessThan, num)

		luckyDog.setRange(1, 10).Run(num)
		t.Log(luckyDog.users)
		So(len(luckyDog.users), ShouldEqual, num)

		luckyDog.setRange(1, 100000).Run(num)
		t.Log(luckyDog.users)
		So(len(luckyDog.users), ShouldEqual, num)
	})
}
