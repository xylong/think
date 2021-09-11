package goroutine

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetUser(t *testing.T) {
	Convey("测试协程池", t, func() {
		p := GetUser()
		user := p.Get().(*user)
		t.Log(user)
		user.name = "jj"
		p.Put(user)
		user2 := p.Get()
		t.Log(user2)
	})
}

func ExampleGetUser() {
	p := GetUser()
	user := p.Get().(*user)
	fmt.Println(user)
	user.name = "jj"
	p.Put(user)
	user2 := p.Get()
	fmt.Println(user2)
}
