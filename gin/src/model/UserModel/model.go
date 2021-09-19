package UserModel

// User 用户
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// NewUser 创建用户模型
func New(attrs ...Attr) *User {
	user := &User{}
	Attrs(attrs).Apply(user)
	return user
}
