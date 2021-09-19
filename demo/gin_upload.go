package demo

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

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
		fileSuffix := u.getSuffix(header.Filename)
		filePrefix := u.getPrefix(header.Filename, fileSuffix)

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

			u.saveBlock(fmt.Sprintf("%s_%d%s", filePrefix, index, fileSuffix), buf)
			index++
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	r.GET("file", func(c *gin.Context) {
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		c.Writer.Header().Set("Content-type", "image/png")

		fileName := c.Query("name")
		fileSuffix := u.getSuffix(fileName)
		filePrefix := u.getPrefix(fileName, fileSuffix)

		for i := 0; i <= 5; i++ {
			f, _ := os.Open(fmt.Sprintf("../storage/%s_%d%s", filePrefix, i, fileSuffix))
			b, _ := ioutil.ReadAll(f)
			c.Writer.Write(b)
			c.Writer.(http.Flusher).Flush()
		}
	})

	r.Run()
}

// 分片存储
func (u *upload) saveBlock(name string, buf []byte) {
	file, err := os.OpenFile("../storage/"+name, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write(buf)
}

func (u *upload) getPrefix(name, suffix string) string {
	return strings.TrimSuffix(path.Base(name), suffix)
}

func (u *upload) getSuffix(name string) string {
	return path.Ext(name)
}
