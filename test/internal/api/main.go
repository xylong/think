package main

import (
	"think/gin/src/api"
	"think/gin/src/db"
	"think/gin/src/handler"
	"think/gin/src/middleware"

	_ "think/gin/src/validator"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	db.InitTable()

	r := gin.New()
	r.Use(middleware.ErrorHandler())

	r.GET("users", api.Index)
	r.GET("users/:id", api.Show)
	r.POST("users", api.Store)
	r.PATCH("users", api.Update)

	//! 统一封装json返回例子
	r.GET("example", handler.Handle()(api.Example))

	r.POST("upload", handler.Handle()(api.Upload))
	r.GET("file")

	r.Run()
}
