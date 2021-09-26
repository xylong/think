package api

import (
	"fmt"
	"io"
	"os"
	"think/gin/src/handler"

	"github.com/gin-gonic/gin"
)

// Upload 上传
func Upload(c *gin.Context) (int, string, interface{}) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		panic(err)
	}

	block := header.Size / 5
	fileSuffix := handler.GetSuffix(header.Filename)
	filePrefix := handler.GetPrefix(header.Filename, fileSuffix)

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

		saveBlock(fmt.Sprintf("%s_%d%s", filePrefix, index, fileSuffix), buf)
		index++
	}

	return 0, "ok", nil
}

func saveBlock(name string, buf []byte) {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write(buf)
}
