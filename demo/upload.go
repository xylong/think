package demo

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

const (
	Format1 = "\b\b%d%%"
	Format2 = "\b\b\b%d%%"
	Format3 = "\b\b\b%d%%\b"
)

// Uploader 上传
type Uploader interface {
	// 上传
	Upload(string) string
}

// SingleUpload 但文件上传
type SingleUpload struct {
	file     *os.File
	complete chan int64
}

func NewSingleUpload() *SingleUpload {
	return &SingleUpload{}
}

// Upload 上传
func (su *SingleUpload) Upload(fileName string) string {
	if err := su.open(fileName); err != nil {
		log.Fatal(err)
		return ""
	}

	newName := su.saveName(fileName)
	ioutil.WriteFile(newName, su.readFile(false), 0600)
	return newName
}

// Uploading 上传显示进度
func (su *SingleUpload) Uploading(fileName string) string {
	if err := su.open(fileName); err != nil {
		log.Fatal(err)
	}

	newName := su.saveName(fileName)
	ioutil.WriteFile(newName, su.readFile(true), 0600)
	return newName
}

func (su *SingleUpload) open(fileName string) (err error) {
	su.file, err = os.Open(fileName)
	return
}

// 读取文件
func (su *SingleUpload) readFile(flag bool) []byte {
	var (
		err   error
		bytes []byte
		info  os.FileInfo
	)

	if flag {
		su.complete = make(chan int64)
		info, err = su.file.Stat()
		if err != nil {
			log.Fatal(err)
			return nil
		}

		fmt.Print("complete:0%")
		go su.Process()
	}

	for {
		buf := make([]byte, 1024*100)

		n, err := su.file.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		time.Sleep(time.Millisecond * 100)
		bytes = append(bytes, buf...)

		if flag {
			go func() {
				su.complete <- int64(len(bytes) * 100 / int(info.Size()))
			}()
		}
	}

	return bytes
}

func (su *SingleUpload) saveName(fileName string) string {
	return su.name() + su.suffix(fileName)
}

func (su *SingleUpload) name() string {
	return time.Now().Format("20060102150405")
}

func (su *SingleUpload) suffix(fileName string) string {
	return path.Ext(fileName)
}

// Process 上传进度
func (su *SingleUpload) Process() {
	format := Format1
	var lastNum int64 = 0

	for v := range su.complete {
		if lastNum >= 10 && v > 10 && v < 100 {
			format = Format2
		} else if v >= 100 {
			format = Format3
		}
		fmt.Printf(format, v)
		lastNum = v
	}
}
