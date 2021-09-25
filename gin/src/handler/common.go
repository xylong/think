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

func init() {
	ResultPool = &sync.Pool{
		New: func() interface{} {
			return NewJSONResult(0, "", nil)
		},
	}
}

type (
	// json结果
	JSONResult struct {
		// 状态码
		Code int `json:"code"`
		// 消息
		Message string `json:"message"`
		// 数据
		Data interface{} `json:"data"`
	}

	// 属性方法
	Attr func(*JSONResult)
	// 属性集合
	Attrs []Attr
)

// NewJSONResult 创建json结果实例
func NewJSONResult(code int, message string, data interface{}) *JSONResult {
	return &JSONResult{Code: code, Message: message, Data: data}
}

// Apply 为JSONResult设置属性
func (a Attrs) Apply(result *JSONResult) {
	for _, f := range a {
		f(result)
	}
}

// Code 设置状态码
func Code(code int) Attr {
	return func(j *JSONResult) {
		j.Code = code
	}
}

// Message 设置消息
func Message(message string) Attr {
	return func(j *JSONResult) {
		j.Message = message
	}
}

// Data 设置返回数据
func Data(data interface{}) Attr {
	return func(j *JSONResult) {
		j.Data = data
	}
}

type (
	// 结果
	Result func(attrs ...Attr) func(output Output)
	// 输出
	Output func(*gin.Context, interface{})
)

// R 返回结果，装饰器模式函数
// 传入响应格式函数，根据不同格式返回响应结果数据
func R(c *gin.Context) Result {
	//? 装饰Result，返回Output
	return func(attrs ...Attr) func(output Output) {
		//! 从池中取result实例，用完放回
		result := ResultPool.Get().(*JSONResult)
		defer ResultPool.Put(result)

		// 设置属性
		Attrs(attrs).Apply(result)

		//* 解偶响应，根据传入的响应函数进行调用
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

// Api api函数
// * 返回code，string，data
type Api func(*gin.Context) (int, string, interface{})

// Handle 输出统一json格式
// 装饰器模式封装返回格式为(code,message,data)的api函数
func Handle() func(api Api) gin.HandlerFunc {
	// ? 装饰api函数
	return func(api Api) gin.HandlerFunc {
		// ? 装饰HandlerFunc
		return func(c *gin.Context) {
			// ! 调用HandlerFunc，返回json
			code, message, data := api(c)
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": message,
				"data":    data,
			})
		}
	}
}
