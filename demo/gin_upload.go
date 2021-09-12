package demo

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

type upload struct {
}

func newUpload() *upload {
	return &upload{}
}

func (u *upload) Run() {
	r := gin.New()

	// 中间件获取错误
	r.Use(func(c *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": e,
				})
			}
		}()
		c.Next()
	})

	r.POST("upload", func(ctx *gin.Context) {
		file, header, err := ctx.Request.FormFile("file")
		if err != nil {
			panic(err)
		}

		block := header.Size / 5

		index := 0
		for {
			buf := make([]byte, block)
			n, err := file.Read(buf)
			if err != nil && err != io.EOF {
				panic(err.Error())
			}
			if n == 0 {
				break
			}

			u.saveBlock(fmt.Sprintf("%s_%d%s", header.Filename, index, path.Ext(header.Filename)), buf)
			index++
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	r.Run()
}

// 分片存储
func (u *upload) saveBlock(name string, buf []byte) {
	save, err := os.OpenFile("../storage/"+name, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	defer save.Close()
	save.Write(buf)
}
