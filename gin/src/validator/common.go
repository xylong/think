package validator

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var valid *validator.Validate

func init() {
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
