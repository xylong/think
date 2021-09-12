package demo

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUpload_Run(t *testing.T) {
	convey.Convey("分片存储，分块显示", t, func() {
		newUpload().Run()
	})
}
