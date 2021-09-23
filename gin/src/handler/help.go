package handler

import (
	"crypto/md5"
	"fmt"
)

// Md5Encrypt md5加密
func Md5Encrypt(s string) string {
	if len(s) != 0 {
		return fmt.Sprintf("%x", md5.Sum([]byte(s)))
	}
	return ""
}
