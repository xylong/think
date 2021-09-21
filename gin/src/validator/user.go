package validator

import (
	"github.com/go-playground/validator/v10"
)

func init() {
	registerValidation("UserName", UserName("required,min=2,max=10").toFunc())
}

// UserName 用户名验证规则
type UserName string

func (un UserName) toFunc() validator.Func {
	return func(fl validator.FieldLevel) bool {
		v, ok := fl.Field().Interface().(string)
		if ok {
			return un.validate(v)
		}

		return false
	}
}

func (un UserName) validate(v string) bool {
	if err := valid.Var(v, string(un)); err != nil {
		return false
	}
	return true
}
