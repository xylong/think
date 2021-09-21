package mapper

import (
	"think/gin/src/model/UserModel"
	"time"

	"github.com/Masterminds/squirrel"
)

type UserMapper struct {
}

// GetUserList 获取用户列表
func (um *UserMapper) GetUserList() *SqlMapper {
	return Mapper(squirrel.Select("id", "name", "gender").From("users").OrderBy("id desc").Limit(10).ToSql())
}

// GetUserByID 根据🆔查询用户
func (um *UserMapper) GetUserByID(id int) *SqlMapper {
	return Mapper(squirrel.Select("id", "name", "gender").Where("id=?", id).From("users").OrderBy("id desc").ToSql())
}

// CreateUser 创建用户
func (um *UserMapper) CreateUser(user *UserModel.User) *SqlMapper {
	now := time.Now()
	return Mapper(squirrel.Insert(user.TableName()).
		Columns("name", "password", "gender", "created_at", "updated_at").
		Values(user.Name, user.Password, user.Gender, now, now).
		ToSql())
}
