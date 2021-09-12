package main

import "think/demo"

func main() {
	upload := demo.NewSingleUpload()
	upload.Uploading("public/婚礼.mp4")
}
