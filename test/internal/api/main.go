package main

import (
	"net/http"
	"think/gin/src/middleware"
	"think/gin/src/model/UserModel"
	"think/gin/src/result"

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

	r.POST("user", func(c *gin.Context) {
		user := UserModel.New()
		result.Result(c.ShouldBind(user)).Unwrap()
		c.JSON(http.StatusOK, user)
	})

	r.Run()
}
