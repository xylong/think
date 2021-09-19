package UserModel

type (
	// 属性方法
	Attr func(*User)
	// Attributes 属性集合
	Attrs []Attr
)

// WithID 设置🆔
func WithID(id int) Attr {
	return func(u *User) {
		u.ID = id
	}
}

// WithName 设置名字
func WithName(name string) Attr {
	return func(u *User) {
		u.Name = name
	}
}

// Apply 应用属性
// 初始化实力时调用
func (a Attrs) Apply(u *User) {
	for _, f := range a {
		f(u)
	}
}
