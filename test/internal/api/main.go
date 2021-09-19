package main

import (
	"net/http"
	"think/gin/src/model/UserModel"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		user := UserModel.New(
			UserModel.WithID(1),
			UserModel.WithName("静静"),
		)
		c.JSON(http.StatusOK, user)
	})

	r.Run()
}
