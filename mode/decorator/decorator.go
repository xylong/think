package decorator

import (
	"fmt"
	"time"
)

func demo(i int) int {
	return i * 10
}

// 被装饰函数签名
type demoFunc func(int) int

// DemoConsume demo耗时
// ? f 函数签名
// * 装饰demo函数，记录demo耗时
func DemoConsume(f demoFunc) demoFunc {
	return func(i int) int {
		start := time.Now()
		result := f(i)
		end := time.Now()

		defer func() {
			fmt.Println(end.Sub(start))
		}()

		return result
	}
}
