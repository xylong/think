package handler

import (
	"fmt"
	"think/gin/src/model/UserModel"
	"think/gin/src/result"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	user := UserModel.New()
	result.Result(c.ShouldBind(user)).Unwrap()
	OK(c)(0, "", result.Result(getInfo(user.ID)).Unwrap())
}

func getInfo(id int) (gin.H, error) {
	if id > 0 {
		return gin.H{"message": "test"}, nil
	} else {
		return nil, fmt.Errorf("test error")
	}
}
