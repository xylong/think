package validator

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var (
	// 自定义验证器
	valid *validator.Validate
	// 验证提示
	tip map[string]string
)

func init() {
	tip = make(map[string]string)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		valid = v
	} else {
		log.Fatalln("error validator")
	}
}

// 注册自定义验证规则
// tag 规则名称
// fn 规则方法
func registerValidation(tag string, fn validator.Func) {
	if err := valid.RegisterValidation(tag, fn); err != nil {
		log.Fatalln(fmt.Sprintf("validator %s error", tag))
	}
}

// IsValidateError 判断是否为验证错误，如果是则抛出错误信息
func IsValidateError(errors error) {
	// 判断错误是否为验证错误
	if errs, ok := errors.(validator.ValidationErrors); ok {
		for _, err := range errs {
			// 判断是否有自定义错误信息，有就抛出
			// todo 处理自带验证报错信息
			if msg, ok := tip[err.Tag()]; ok {
				panic(msg)
			}
		}
	}
}
