package model

// 创建动物
// * Attr 属性
type AnimalCreate func(...Attr) interface{}

// 属性设置函数
type Attr func(interface{})

// 属性方法集合
type Attrs []Attr

// Apply 设置模型属性
// ! m 具体模型
func (this Attrs) Apply(m interface{}) {
	for _, f := range this {
		f(m)
	}
}
