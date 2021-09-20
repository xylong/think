package UserModel

import (
	"time"
)

// User 用户
type User struct {
	ID        int       `gorm:"column:id;type:int(11);primaryKey;autoIncrement;" json:"id" form:"id"`
	Name      string    `gorm:"column:name;type:varchar(20);unique;not null;comment:'名字'" json:"name" form:"name" binding:"min=2"`
	Password  string    `gorm:"column:password;type:char(32);not null;comment:'密码'" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at;not null;comment:'创建时间'" json:"created_at"`
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
