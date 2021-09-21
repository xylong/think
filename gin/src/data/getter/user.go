package getter

import (
	"fmt"
	"think/gin/src/data/mapper"
	"think/gin/src/model/UserModel"
	"think/gin/src/result"
)

var UserGetter IUserGetter

func init() {
	UserGetter = NewUserGetterImpl()
}

type IUserGetter interface {
	// 用户列表
	GetUserList() []*UserModel.User
	// 用户详情
	GetUserByID(id int) *result.Error
}

// UserGetterImpl 用户获取器
type UserGetterImpl struct {
	mapper *mapper.UserMapper
}

func NewUserGetterImpl() *UserGetterImpl {
	return &UserGetterImpl{
		mapper: &mapper.UserMapper{},
	}
}

// GetUserList 获取用户列表
func (u *UserGetterImpl) GetUserList() (users []*UserModel.User) {
	u.mapper.GetUserList().Query().Find(&users)
	return
}

// GetUserByID 根据🆔获取用户
func (u *UserGetterImpl) GetUserByID(id int) *result.Error {
	user := UserModel.New()
	r := u.mapper.GetUserByID(id).Query().Find(user)
	if r.Error != nil || r.RowsAffected == 0 {
		return result.Result(nil, fmt.Errorf("not found user,id=%d", id))
	}
	return result.Result(user, nil)
}
