package basic

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLuckyUser_Run(t *testing.T) {
	Convey("抽取幸运观众", t, func() {
		num := 10
		luckyDog := NewLuckyUser()

		luckyDog.setRange(1, 5).Run(num)
		ShouldBeLessThan(len(luckyDog.users), num)
		t.Log(luckyDog.users)

		luckyDog.setRange(1, 10).Run(num)
		ShouldEqual(len(luckyDog.users), num)
		t.Log(luckyDog.users)

		luckyDog.setRange(1, 100000).Run(num)
		ShouldEqual(len(luckyDog.users), num)
		t.Log(luckyDog.users)
	})
}
