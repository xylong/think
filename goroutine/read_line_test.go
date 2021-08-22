package goroutine

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestReader_Read(t *testing.T) {
	Convey("逐行读取文件", t, func() {
		newReader().Read()
	})
}
