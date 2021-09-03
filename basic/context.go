package basic

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// LuckyUser 幸运用户
type LuckyUser struct {
	min, max int
	users    []int
	wg       sync.WaitGroup
}

func NewLuckyUser() *LuckyUser {
	return &LuckyUser{}
}

// 设置取值范围
func (l *LuckyUser) setRange(min, max int) *LuckyUser {
	l.min, l.max = min, max
	return l
}

// 抽取用户
func (l *LuckyUser) draw() int {
	return rand.Intn(l.max-l.min) + l.min
}

// 获取用户
func (l *LuckyUser) get(ctx context.Context, num int) {
	sum, id := 0, 0

loop:
	for {
		select {
		// 父context发起取消操作
		case <-ctx.Done():
			l.wg.Done()
			fmt.Println(sum)
			break loop
		default:
			if len(l.users) < num {
				id = l.draw()

				if l.find(id) == -1 {
					l.users = append(l.users, id)
				}
			}
			sum++
		}
	}
}

// 查找切片中是否存在指定值
// val指定查找的元素
func (l *LuckyUser) find(val int) int {
	for index, item := range l.users {
		if item == val {
			return index
		}
	}
	return -1
}

// Run 运行
// num 限制人数
func (l *LuckyUser) Run(num int) {
	rand.Seed(time.Now().Unix())
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)

	l.wg.Add(1)
	go l.get(ctx, num)
	l.wg.Wait()
}
