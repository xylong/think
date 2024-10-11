package main

import (
	"fmt"
	"strconv"
)

// 位运算实现权限组合
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

// showPermission 查看有哪些权限
func showPermission(role int64) []int64 {
	var (
		arr = []int64{Read, Create, Update, Delete}
		brr []int64
	)

	for _, i := range arr {
		if hasPermission(role, i) {
			brr = append(brr, i)
		}
	}

	return brr
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

	// 查看有哪些权限
	fmt.Println(showPermission(Writer))
}
