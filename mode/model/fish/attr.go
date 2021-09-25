package fish

import "think/mode/model"

// ID 设置id
func ID(id int) model.Attr {
	return func(m interface{}) {
		m.(*Fish).ID = id
	}
}

// Name 设置名字
func Name(name string) model.Attr {
	return func(m interface{}) {
		m.(*Fish).Name = name
	}
}
