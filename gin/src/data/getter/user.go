package getter

import (
	"fmt"
	"think/gin/src/db"
	"think/gin/src/model/UserModel"
	"think/gin/src/result"
)

var UserGetter IUserGetter

func init() {
	UserGetter = NewUserGetterImpl()
}

type IUserGetter interface {
	GetUserList() []*UserModel.User
	GetUserByID(id int) *result.Error
}

type UserGetterImpl struct{}

func NewUserGetterImpl() *UserGetterImpl {
	return &UserGetterImpl{}
}

func (u *UserGetterImpl) GetUserList() (users []*UserModel.User) {
	db.Orm.Find(&users)
	return
}

func (u *UserGetterImpl) GetUserByID(id int) *result.Error {
	user := UserModel.New()
	r := db.Orm.Where("id=?", id).Find(user)
	if r.Error != nil || r.RowsAffected == 0 {
		return result.Result(nil, fmt.Errorf("not found user,id=%d", id))
	}
	return result.Result(user, nil)
}
