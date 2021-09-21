package mapper

import (
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
