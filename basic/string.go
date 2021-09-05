package basic

import "fmt"

// String 自定义字符串
type String string

// Length 字符串长度
func (s String) Length() int {
	return len(s)
}

// Each 遍历字符串
// f 回调函数
func (s String) Each(f func(item string)) {
	for _, item := range s {
		f(fmt.Sprintf("%c", item))
	}
}

// Reverse 反转字符串
func (s String) Reverse() string {
	str := []rune(s)
	max := len(str) - 1

	for i, j := 0, max; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}

	return string(str)
}

// From 字符串转自定义字符串
func From(str string) String {
	return String(str)
}

// FromInt int转自定义字符串
func FromInt(i int) String {
	return String(fmt.Sprintf("%d", i))
}
