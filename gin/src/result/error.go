package result

import (
	"fmt"
	"think/gin/src/validator"
)

//? ErrorResult 有error的返回结果
type ErrorResult struct {
	data interface{}
	err  error
}

// Unwrap 判断error不为空，panic error
func (e *ErrorResult) Unwrap() interface{} {
	if e.err != nil {
		validator.IsValidateError(e.err)
		panic(e.err.Error())
	}

	return e.data
}

// UnwrapOr 判断error是否为空，返回自定义参数
func (e *ErrorResult) UnwrapOr(any interface{}) interface{} {
	if e.err != nil {
		return any
	}

	return e.data 
}

// UnwrapFun 判断error不为空，调用自定义函数
func (e *ErrorResult) UnwrapFun(f func () interface{}) interface{} {
	if e.err != nil {
		return f()
	}

	return e.data
}

// Result 响应结果
func Result(any ...interface{}) *ErrorResult {
	if len(any) == 1 {
		if any[0] == nil {
			return &ErrorResult{nil, nil}
		}
		if err, ok := any[0].(error); ok {
			return &ErrorResult{nil, err}
		}
	}

	if len(any) == 2 {
		if any[1] == nil {
			return &ErrorResult{any[0], nil}
		}
		if err, ok := any[1].(error); ok {
			return &ErrorResult{any[0], err}
		}
	}

	return &ErrorResult{nil, fmt.Errorf("error result")}
}
