package main

import (
	"think/gin/src/controller"
	"think/gin/src/db"
	"think/gin/src/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	db.InitTable()

	r := gin.New()
	r.Use(middleware.ErrorHandler())

	r.GET("users", controller.Index)
	r.GET("users/:id", controller.Show)
	r.POST("users", controller.Store)

	r.Run()
}
