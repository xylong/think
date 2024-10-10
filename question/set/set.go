package main

import (
	"bytes"
	"fmt"
)

/*
q：用go实现一个简单set？
思路：用map的key来存储元素，因为mao的key是不重复的
*/

// Empty 空值
// 用作map的value
type Empty struct{}

type Set map[any]Empty

// Add 往set添加元素
func (s Set) Add(element ...any) Set {
	for _, e := range element {
		s[e] = Empty{}
	}

	return s
}

// String 字符串
// print打印时自动调用
func (s Set) String() string {
	var buff bytes.Buffer

	for key := range s {
		if buff.Len() > 0 {
			buff.WriteString(",")
		}
		buff.WriteString(fmt.Sprintf("%v", key))
	}

	return buff.String()
}

func NewSet() Set {
	return make(map[any]Empty)
}

func main() {
	set := NewSet()
	set.Add(1, 2, 3, 2, 1, "a", "b", "a", "b", "c")
	fmt.Println(set)
}
