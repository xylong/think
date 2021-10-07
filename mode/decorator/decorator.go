package decorator

import (
	"fmt"
	"time"
)

// Sum 累加求和，从a加到b
func Sum(a, b int) int {
	result := 0

	for i := a; i <= b; i++ {
		result = result + i
	}

	return result
}

// MathFunc 数学计算
type MathFunc func(a, b int) int

type GoMathFunc func(int) int

// Consume 计算耗时
// 装饰器装饰数学计算函数
func Consume(mf MathFunc) MathFunc {
	return func(i1, i2 int) int {
		start := time.Now()
		result := mf(i1, i2)
		end := time.Now()

		defer func() {
			fmt.Println(end.Sub(start))
		}()

		return result
	}
}
