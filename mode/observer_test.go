package mode

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

type (
	// 商品
	goods struct {
		id    int
		name  string
		price float64
		stock int
	}
	// 日志
	log struct{}
	// 订单
	order struct {
		no        string //订单号
		good      *goods
		observers []observer // 观察者
	}
)

// notify 通知观察者
func (o *order) notify() {
	for _, observer := range o.observers {
		observer.handle(o)
	}
}

// add 添加观察者
func (o *order) add(observer observer) {
	o.observers = append(o.observers, observer)
}

// create 创建订单
func (o *order) create(g *goods) {
	o.no = fmt.Sprintf("%s%d", time.Now().Format("20060102150405"), g.id) // 生成订单号
	o.good = g
	o.notify()
}

func (g *goods) handle(obj interface{}) {
	g.stock--
}

func (l *log) handle(obj interface{}) {
	order := obj.(*order)
	fmt.Printf("订单号：%s，商品：%s剩余%d\n", order.no, order.good.name, order.good.stock)
}

func TestObserver(t *testing.T) {
	Convey("观察者模式：商品和日志观察订单变化修改库存并记录日志", t, func() {
		stock := 10

		// 观察者
		g, l := &goods{12, "book", 99.8, stock}, &log{}
		// 被观察者
		o := &order{
			observers: []observer{g},
		}
		o.add(l)

		o.create(g)
		So(g.stock, ShouldEqual, stock-1)
	})
}
