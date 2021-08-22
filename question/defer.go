package question

import "fmt"

// Defer 延迟执行
type Defer struct {
}

func NewDefer() *Defer {
	return &Defer{}
}

// DefinedParam 定义了参数
// 如果defer执行有参数，则定义时就已经确定，不会再改变
// 输出1
func (d *Defer) DefinedParam() {
	a := 1
	defer fmt.Println(a)
	a++
}

// UndefinedParam 未定义参数
// 输出2
func (d *Defer) UndefinedParam() {
	a := 1
	defer func() {
		fmt.Println(a)
	}()
	a++
}

// BreakThrough 突破defer默认参数机制
// 突破DefinedParam机制
func (d *Defer) BreakThrough() {
	a := 1
	defer d.show(&a)
	a++
}

func (d *Defer) show(i *int) {
	fmt.Println(*i)
}

// Echo 输出
// defer链式调用
func (d *Defer) Echo(any interface{}) *Defer {
	fmt.Println(any)
	return d
}

// Loop 循环
// defer循环调用
func (d *Defer) Loop() {
	// 3,3,3
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
	// 2,1,0,后进先出
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}

// Panic 异常
// defer>panic,321ab
func (d *Defer) Panic() {
	defer func() {
		defer fmt.Println(1)
		defer fmt.Println(2)
		defer fmt.Println(3)

		panic("b")
	}()

	panic("a")
}
