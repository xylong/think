package factory

import (
	"time"
)

// Demo 演示
type Demo struct {
	// 开始时间
	start time.Time

	// 结束时间
	end time.Time
}

// Start 设置开始时间
func (d *Demo) Start() {
	d.start = time.Now()
}

// End 设置结束时间
func (d *Demo) End() {
	d.end = time.Now()
}

func (d *Demo) TakeUpTime() time.Duration {
	return d.end.Sub(d.start)
}
