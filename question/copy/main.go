package main

import "fmt"

/*
浅拷贝是对对象的引用（地址）的复制，新旧对象还是指向同一个内存地址
深拷贝是对对象的全复制，包括对象内部的所有引用的对象
*/

// 深拷贝
func deepCopy() {
	arr := []int{1, 2, 3}
	brr := make([]int, len(arr), cap(arr))
	copy(brr, arr)

	arr[0] = 100
	fmt.Println(arr, brr)
}

// 浅拷贝
// arr和brr指针都指向同一个底层数组，一个改了另外一个也会跟着变
func shallowCopy() {
	arr := []int{1, 2, 3}
	brr := arr

	arr[0] = 100
	fmt.Println(arr, brr)
}

func main() {
	shallowCopy()
	fmt.Println("-----------")
	deepCopy()
}
