package main

import (
	"think/gin/src/controller"
	"think/gin/src/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(middleware.ErrorHandler())

	r.GET("users", controller.Index)
	r.POST("users", controller.Store)

	r.Run()
}
