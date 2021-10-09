package goroutine

import (
	"math/rand"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMyDogStruct_Add(t *testing.T) {
	Convey("协程和chanel传值", t, func() {
		task:=NewTask()

		for i := 0; i < 100; i++ {
			task.Add(func() interface{} {
				return rand.Intn(100)
			})
		}

		task.Range(func(i interface{}) {
			t.Log(i)
		})
	})
}
