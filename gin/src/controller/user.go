package controller

import (
	"think/gin/src/data/getter"
	. "think/gin/src/handler"
	"think/gin/src/model/UserModel"
	"think/gin/src/result"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	R(c)(Data(getter.UserGetter.GetUserList()))(OK)
}

func Show(c *gin.Context) {
	id := &struct {
		ID int `uri:"id" binding:"required,gt=0"`
	}{}

	result.Result(c.ShouldBindUri(id)).Unwrap()
	R(c)(Data(getter.UserGetter.GetUserByID(id.ID).Unwrap()))(OK)
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
