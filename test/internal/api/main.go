package main

import (
	"net/http"
	"think/gin/src/handler"
	"think/gin/src/middleware"
	"think/gin/src/model/UserModel"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(middleware.ErrorHandler())

	r.GET("/", func(c *gin.Context) {
		user := UserModel.New().Mutate(
			UserModel.WithID(1),
			UserModel.WithName("静静"),
		)
		c.JSON(http.StatusOK, user)
	})

	r.POST("users", handler.UserList)

	r.Run()
}
