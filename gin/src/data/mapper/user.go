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
