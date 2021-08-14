package main

import (
	. "think/iface"
)

func main() {
	demo := NewDemoFactory().Create("sum")
	//demo.Run()
	//demo.TakeUpTime()
	demo.Go()
	demo.TakeUpTime()
}
