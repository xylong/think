package proxy

import "fmt"

type (
	// Login 登录方法
	Login func(name, password string)

	// LoginLog 登录日志,装饰登录，记录日志
	LoginLog func(Login) Login
)

// Log2Mongo 记录mongodb日志
func Log2Mongo(l Login) Login {
	return func(name, password string) {
		l(name, password)
		fmt.Println("记录到mongodb")
	}
}

// Log2Mysql 记录mysql日志
func Log2Mysql(l Login) Login {
	return func(name, password string) {
		l(name, password)
		fmt.Println("记录到mysql")
	}
}

// UserService 用户服务
type UserService struct{}

// Login 登录
func (us *UserService) Login(name, password string) {
	fmt.Println("登录成功")
}

// UserProxy 代理用户服务
type UserProxy struct {
	service *UserService
}

func NewUserProxy(service *UserService) *UserProxy {
	return &UserProxy{service: service}
}

// Login 代理登录
func (us *UserProxy) Login(name, password string) {
	fmt.Println("记录日志")
	us.service.Login(name, password)
}

// Login2 代理+装饰器
func (us *UserProxy) Login2(ll LoginLog) Login {
	return ll(us.Login)
}
