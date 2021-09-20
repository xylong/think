package handler

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// ResultPool 结果池
// 从池中取json结果实例，避免重复创建
var ResultPool *sync.Pool

type (
	// json结果
	JSONResult struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	// 结果
	Result func(code int, message string, data interface{})
)

func init() {
	ResultPool = &sync.Pool{
		New: func() interface{} {
			return NewJSONResult(0, "", nil)
		},
	}
}

func NewJSONResult(code int, message string, data interface{}) *JSONResult {
	return &JSONResult{Code: code, Message: message, Data: data}
}

// OK 成功
// 装饰器模式封装成功返回结果
func OK(c *gin.Context) Result {
	return func(code int, message string, data interface{}) {
		// 从池中取结果实例，用完放回
		result := ResultPool.Get().(*JSONResult)
		defer ResultPool.Put(result)

		result.Code = code
		result.Message = message
		result.Data = data

		c.JSON(http.StatusOK, result)
	}
}
