package validator

import (
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
