package demo

import "testing"

func TestSingleUpload_Upload(t *testing.T) {
	upload := NewSingleUpload()
	upload.Upload()
}
