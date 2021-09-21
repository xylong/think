package validator

import (
	"log"

	"github.com/go-playground/validator/v10"
)

func init() {
	if err := valid.RegisterValidation("UserName", VUserName); err != nil {
		log.Fatalln("validator UserName error")
	}
}

var VUserName validator.Func = func(fl validator.FieldLevel) bool {
	name, ok := fl.Field().Interface().(string)
	length := len([]rune(name))

	if ok && length >= 2 && length <= 10 {
		return true
	}
	return false
}
