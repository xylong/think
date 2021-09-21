package UserModel

import (
	"gorm.io/gorm"
)

// User 用户
type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;" json:"id" form:"id"`
	Name     string `gorm:"type:varchar(20);unique;not null;comment:'名字'" json:"name" form:"name" binding:"UserName"`
	Password string `gorm:"type:char(32);not null;comment:'密码'" json:"password" form:"password" binding:"required,min=6"`
	Gender   int    `gorm:"type:tinyint(1);not null;default:0;comment:'性别0女 1男'" json:"gender"`
	gorm.Model
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
