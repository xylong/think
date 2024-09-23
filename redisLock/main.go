package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"think/redisLock/lib"
	"time"
)

var (
	n = 1
)

func main() {
	r := gin.New()

	r.Use(func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err})
			}
		}()

		ctx.Next()
	})

	r.GET("/", func(ctx *gin.Context) {
		locker := lib.NewLockerWithTTL("lock1", time.Second*5).Lock()
		defer locker.Unlock()

		if ctx.Query("t") != "" {
			time.Sleep(time.Second * 10) // 模拟卡顿
		}

		n++
		ctx.JSONP(http.StatusOK, gin.H{"msg": n})
	})

	r.Run(":8080")
}
