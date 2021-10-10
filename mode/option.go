package mode

type (
	// Attrs 属性集合
	Attrs []Attr

	// Attr 属性
	Attr func(interface{})

	// NewEntity 创建实体
	NewEntity func(...Attr) interface{}
)

// Set 设置属性
func (a Attrs) Set(obj interface{}) {
	for _, f := range a {
		f(obj)
	}
}
