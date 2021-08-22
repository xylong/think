package main

import "fmt"

func main() {
	defer func() {
		defer fmt.Println(1)
		defer fmt.Println(2)
		defer fmt.Println(3)

		panic("b")
	}()

	panic("a")
}
