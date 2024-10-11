package main

import (
	"fmt"
	"strconv"
)

const (
	Read   int64 = 1 // 0001
	Create int64 = 2 // 0010
	Update int64 = 4 // 0100
	Delete int64 = 8 // 1000
)

// 判断是否有某个权限
func hasPermission(role, permission int64) bool {
	return role&permission == permission
}

func main() {
	// 添加写操作权限
	var Writer = Create | Update | Delete
	fmt.Println(Writer)                       // 打印10进制
	fmt.Println(strconv.FormatInt(Writer, 2)) // 打印2进制

	fmt.Printf("读权限：%v\n", hasPermission(Writer, Read))    // 没有读权限
	fmt.Printf("删除权限：%v\n", hasPermission(Writer, Delete)) // 有删除权限

	// 去掉删除权限
	Writer = Writer ^ Delete
	fmt.Println(Writer, strconv.FormatInt(Writer, 2))
	fmt.Printf("删除权限：%v\n", hasPermission(Writer, Delete))
}
