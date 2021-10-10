package mode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAttrs_Set(t *testing.T) {
	Convey("选项模式", t, func() {
		type user struct {
			name   string
			gender int
			age    int
		}

		newUser := func(attrs ...Attr) *user {
			user := &user{}
			Attrs(attrs).Set(user)
			return user
		}

		withName := func(name string) Attr {
			return func(i interface{}) {
				i.(*user).name = name
			}
		}

		withGender := func(gender int) Attr {
			return func(i interface{}) {
				i.(*user).gender = gender
			}
		}

		withAge := func(age int) Attr {
			return func(i interface{}) {
				i.(*user).age = age
			}
		}

		user1 := newUser(withName("小明"), withGender(1), withAge(18))
		So(user1.name,ShouldEqual,"小明")
		So(user1.gender,ShouldEqual,1)
		So(user1.age,ShouldEqual,18)
		user2 := newUser(withName("小红"), withGender(0))
		So(user2.name,ShouldEqual,"小红")
		So(user2.gender,ShouldEqual,0)
		So(user2.age,ShouldEqual,0)

		t.Log(user1,user2)
	})
}
