package api

import (
	"think/gin/src/data/getter"
	"think/gin/src/data/setter"
	. "think/gin/src/handler"
	"think/gin/src/model/UserModel"
	"think/gin/src/result"

	"github.com/gin-gonic/gin"
)

//? 用户列表
func Index(c *gin.Context) {
	R(c)(Data(getter.UserGetter.GetUserList()))(OK)
}

//？ 用户详情
func Show(c *gin.Context) {
	id := &struct {
		ID int `uri:"id" binding:"required,gt=0"`
	}{}

	result.Result(c.ShouldBindUri(id)).Unwrap()
	R(c)(Data(getter.UserGetter.GetUserByID(id.ID).Unwrap()))(OK)
}

//? 创建用户
func Store(c *gin.Context) {
	user := UserModel.New()
	result.Result(c.ShouldBindJSON(user)).Unwrap()
	R(c)(Data(setter.UserSetter.CreateUser(user).Unwrap()))(OK)
}

//? 更新用户
func Update(c *gin.Context) {
	user := UserModel.New()
	result.Result(c.ShouldBindJSON(user)).Unwrap()
	R(c)(Data(setter.UserSetter.UpdateUser(user).Unwrap()))(OK)
}

//? Example 统一输出json例子
func Example(c *gin.Context) (int, string, interface{}) {
	return 0, "", nil
}
