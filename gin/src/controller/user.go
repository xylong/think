package controller

import (
	. "think/gin/src/handler"
	"think/gin/src/model/UserModel"
	"think/gin/src/result"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	user := UserModel.New().Mutate(
		UserModel.WithID(1),
		UserModel.WithName("静静"),
	)
	R(c)(Data(user))(OK)
}

func Store(c *gin.Context) {
	user := UserModel.New()
	result.Result(c.ShouldBind(user)).Unwrap()

	if user.ID > 0 {
		R(c)(Data("user list"))(OK)
	} else {
		R(c)(
			Code(10001),
			Message("用户错误"),
		)(Error)
	}
}
