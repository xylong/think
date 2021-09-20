package handler

import (
	"think/gin/src/model/UserModel"
	"think/gin/src/result"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	user := UserModel.New()
	result.Result(c.ShouldBind(user)).Unwrap()

	if user.ID > 0 {
		R(c)(0, "", "list")(OK)
	} else {
		R(c)(10001, "用户错误", nil)(Error)
	}
}
