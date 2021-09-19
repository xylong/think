package UserModel

// User 用户
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name" form:"name" binding:"min=2"`
}

// NewUser 创建用户模型
func New(attrs ...Attr) *User {
	user := &User{}
	Attrs(attrs).Apply(user)
	return user
}

// Mutate 设置属性
func (u *User) Mutate(attrs ...Attr) *User {
	Attrs(attrs).Apply(u)
	return u
}
