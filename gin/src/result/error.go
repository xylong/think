package result

// Error 自定义错误
type Error struct {
	err error
}

// Unwrap 展开
func (e *Error) Unwrap() interface{} {
	if e.err != nil {
		panic(e.err.Error())
	}

	return nil
}

// Result 响应结果
func Result(err error) *Error {
	return &Error{err: err}
}
