package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"think/redisLock/lib"
)

func main() {
	r := gin.New()

	r.Use(func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err})
			}
		}()

		ctx.Next()
	})

	r.GET("/", func(ctx *gin.Context) {
		lib.Lock("lock1")
		ctx.JSONP(http.StatusOK, gin.H{"msg": "ok"})
	})

	r.Run(":8080")
}
