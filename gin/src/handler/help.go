package handler

import (
	"crypto/md5"
	"fmt"
	"path"
	"strings"
)

// Md5Encrypt md5加密
func Md5Encrypt(s string) string {
	if len(s) != 0 {
		return fmt.Sprintf("%x", md5.Sum([]byte(s)))
	}
	return ""
}

// GetPrefix 获取文件前缀名
func GetPrefix(name, suffix string) string {
	return strings.TrimSuffix(path.Base(name), suffix)
}

// GetSuffix 获取文件后缀名
func GetSuffix(name string) string {
	return path.Ext(name)
}
