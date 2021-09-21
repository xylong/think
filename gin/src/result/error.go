package result

import (
	"fmt"
	"think/gin/src/validator"
)

// Error 自定义错误
type Error struct {
	data interface{}
	err  error
}

// Unwrap 展开
func (e *Error) Unwrap() interface{} {
	if e.err != nil {
		validator.IsValidateError(e.err)
		panic(e.err.Error())
	}

	return e.data
}

// Result 响应结果
func Result(any ...interface{}) *Error {
	if len(any) == 1 {
		if any[0] == nil {
			return &Error{nil, nil}
		}
		if err, ok := any[0].(error); ok {
			return &Error{nil, err}
		}
	}

	if len(any) == 2 {
		if any[1] == nil {
			return &Error{any[0], nil}
		}
		if err, ok := any[1].(error); ok {
			return &Error{any[0], err}
		}
	}

	return &Error{nil, fmt.Errorf("error result")}
}
