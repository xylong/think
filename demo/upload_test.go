package demo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSingleUpload_Upload(t *testing.T) {
	Convey("单文件上传", t, func() {
		upload := NewSingleUpload()
		t.Log(upload.Upload("../public/婚礼.mp4"))
	})
}

func TestSingleUpload_Uploading(t *testing.T) {
	Convey("显示上传进度", t, func() {
		upload := NewSingleUpload()
		upload.Uploading("../public/婚礼.mp4")
	})
}
