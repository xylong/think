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
	runes := []rune(s)
	max := len(runes) - 1

	for a, b := 0, max; a < b; a, b = a+1, b-1 {
		runes[a], runes[b] = runes[b], runes[a]
	}

	return string(runes)
}

// From 字符串转自定义字符串
func From(str string) String {
	return String(str)
}

// FromInt int转自定义字符串
func FromInt(i int) String {
	return String(fmt.Sprintf("%d", i))
}
