package iface

import (
	"think/goroutine"
)

// IDemo demo
type IDemo interface {
	// Start 运行开始时间
	Start()

	// End 运行结束时间
	End()

	// TakeUpTime 耗时
	TakeUpTime()

	// Run 普通运行
	Run()

	// Go 并发运行
	Go()
}

type DemoFactory struct {
}

func NewDemoFactory() *DemoFactory {
	return &DemoFactory{}
}

func (df *DemoFactory) Create(name string) IDemo {
	switch name {
	case "sum":
		return goroutine.NewSum()
	default:
		return nil
	}
}
