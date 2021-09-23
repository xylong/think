package setter

import (
	"think/gin/src/data/mapper"
	"think/gin/src/model/LogModel"
	"think/gin/src/model/UserModel"
	"think/gin/src/result"
	"time"
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
	userMapper *mapper.UserMapper
	logMapper  *mapper.LogMapper
}

func NewUserSetterImpl() *UserSetterImpl {
	return &UserSetterImpl{
		userMapper: &mapper.UserMapper{},
		logMapper:  &mapper.LogMapper{},
	}
}

// CreateUser 创建用户
func (u *UserSetterImpl) CreateUser(user *UserModel.User) *result.Error {
	r := u.userMapper.CreateUser(user).Exec()
	return result.Result(r.RowsAffected, r.Error)
}

// UpdateUser 更新用户
func (u *UserSetterImpl) UpdateUser(user *UserModel.User) *result.Error {
	// 更新用户
	update := u.userMapper.UpdateUser(user)
	// 记录日志
	addLog := u.logMapper.AddLog(LogModel.New(LogModel.WithName("update user"), LogModel.WithCreatedAt(time.Now())))

	// 事务执行
	err := mapper.Mappers(update, addLog).Exec(func() error {
		if err := update.Exec().Error; err != nil {
			return err
		}
		if err := addLog.Exec().Error; err != nil {
			return err
		}
		return nil
	})

	return result.Result(user, err)
}
