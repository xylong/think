package getter

import (
	"think/gin/src/db"
	"think/gin/src/model/UserModel"
)

var UserGetter IUserGetter

func init() {
	UserGetter = NewUserGetterImpl()
}

type IUserGetter interface {
	GetUserList() []*UserModel.User
}

type UserGetterImpl struct{}

func NewUserGetterImpl() *UserGetterImpl {
	return &UserGetterImpl{}
}

func (u *UserGetterImpl) GetUserList() (users []*UserModel.User) {
	db.Orm.Find(&users)
	return
}
