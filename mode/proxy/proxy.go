package proxy

import "fmt"

// UserService 用户服务
type UserService struct{}

// LoginFunc 登录方法
type LoginFunc func(name, password string)

// LogDecorator 日志装饰器
type LogDecorator func(LoginFunc) LoginFunc

// Log2Mongo 记录mongo日志
func Log2Mongo(lf LoginFunc) LoginFunc {
	return func(name, password string) {
		lf(name, password)
		fmt.Println("日志记录到mongodb")
	}
}

// Log2Mysql 记录mysql日志
func Log2Mysql(lf LoginFunc) LoginFunc {
	return func(name, password string) {
		lf(name, password)
		fmt.Println("日志记录到mysql")
	}
}

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
func (us *UserProxy) Login2(ld LogDecorator) LoginFunc {
	return ld(us.Login)
}
