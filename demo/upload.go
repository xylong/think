package demo

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Uploader 上传
type Uploader interface {
	// 上传
	Upload()
}

// SingleUpload 但文件上传
type SingleUpload struct {
	file *os.File
}

func NewSingleUpload() *SingleUpload {
	return &SingleUpload{}
}

// Upload 上传
func (su *SingleUpload) Upload() {
	file, err := os.Open("金泰熙.gif")
	if err != nil {
		log.Fatalln(err)
	}

	bytes := make([]byte, 0)
	for {
		buf := make([]byte, 1024)

		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		bytes = append(bytes, buf...)
	}

	ioutil.WriteFile("demo.png", bytes, 0600)
}
