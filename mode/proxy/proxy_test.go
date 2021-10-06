package proxy

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUserProxy_Login(t *testing.T) {
	Convey("代理模式", t, func() {
		user := new(UserService)
		userProxy := NewUserProxy(user)
		userProxy.Login("小明", "123456")
	})
}

func TestUserProxy_Login2(t *testing.T) {
	Convey("代理模式+装饰器模式", t, func() {
		user := new(UserService)
		userProxy := NewUserProxy(user)
		userProxy.Login2(Log2Mongo)("小红", "123456")
	})
}
