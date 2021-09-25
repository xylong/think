package bird

import "think/mode/model"

// ID 设置id
func ID(id int) model.Attr {
	return func(m interface{}) {
		m.(*Bird).ID = id
	}
}

// Name 设置名字
func Name(name string) model.Attr {
	return func(m interface{}) {
		m.(*Bird).Name = name
	}
}

// Color 设置颜色
func Color(color string) model.Attr {
	return func(m interface{}) {
		m.(*Bird).Color = color
	}
}
