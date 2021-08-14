package factory

import (
	"time"
)

// Demo 演示
type Demo struct {
	// 开始时间
	StartTime time.Time

	// 结束时间
	EndTime time.Time
}

// Start 设置开始时间
func (d *Demo) Start() {
	d.StartTime = time.Now()
}

// End 设置结束时间
func (d *Demo) End() {
	d.EndTime = time.Now()
}

func (d *Demo) TakeUpTime() time.Duration {
	return d.EndTime.Sub(d.StartTime)
}
