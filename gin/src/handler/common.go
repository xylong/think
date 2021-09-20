package handler

import (
	"fmt"
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
	Result func(code int, message string, data interface{}) func(output Output)
	// 输出
	Output func(*gin.Context, interface{})
)

func init() {
	ResultPool = &sync.Pool{
		New: func() interface{} {
			return NewJSONResult(0, "", nil)
		},
	}
}

// NewJSONResult 创建json结果实例
func NewJSONResult(code int, message string, data interface{}) *JSONResult {
	return &JSONResult{Code: code, Message: message, Data: data}
}

// R 返回结果，装饰器模式函数
// 传入响应格式函数，根据不同格式返回响应结果数据
func R(c *gin.Context) Result {
	// 装饰Result，返回Output
	return func(code int, message string, data interface{}) func(output Output) {
		// 从池中取result实例，用完放回
		result := ResultPool.Get().(*JSONResult)
		defer ResultPool.Put(result)

		// 设置结果属性
		result.Code = code
		result.Message = message
		result.Data = data

		// 解偶响应，根据传入的响应函数进行调用
		return func(output Output) {
			// 调用传入的响应函数，例：OK(c, result)
			output(c, result)
		}
	}
}

// OK 成功
func OK(c *gin.Context, any interface{}) {
	c.JSON(http.StatusOK, any)
}

// Error 失败
func Error(c *gin.Context, any interface{}) {
	c.JSON(http.StatusBadRequest, any)
}

// OK2String 成功，返回字符串
func OK2String(c *gin.Context, any interface{}) {
	c.String(http.StatusOK, fmt.Sprintf("%v", any))
}
