package setter

import (
	"think/gin/src/data/mapper"
	"think/gin/src/model/UserModel"
	"think/gin/src/result"
)

var UserSetter IUserSetter

func init() {
	UserSetter = NewUserSetterImpl()
}

type IUserSetter interface {
	// 创建用户
	CreateUser(*UserModel.User) *result.Error
	// 修改用户
	UpdateUser(*UserModel.User) *result.Error
}

type UserSetterImpl struct {
	mapper *mapper.UserMapper
}

func NewUserSetterImpl() *UserSetterImpl {
	return &UserSetterImpl{
		mapper: &mapper.UserMapper{},
	}
}

// CreateUser 创建用户
func (u *UserSetterImpl) CreateUser(user *UserModel.User) *result.Error {
	r := u.mapper.CreateUser(user).Exec()
	return result.Result(r.RowsAffected, r.Error)
}

// UpdateUser 更新用户
func (u *UserSetterImpl) UpdateUser(user *UserModel.User) *result.Error {
	r := u.mapper.UpdateUser(user).Exec()
	return result.Result(r.RowsAffected, r.Error)
}
