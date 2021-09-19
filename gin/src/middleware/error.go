package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler 错误处理
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err,
				})
			}
		}()

		c.Next()
	}
}
