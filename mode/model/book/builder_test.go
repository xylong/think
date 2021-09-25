package book

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBuild(t *testing.T) {
	Convey("建造者模式", t, func() {
		id, price := 101, 99.0
		book := new(Book).Builder(id).SetPrice(price).Build()
		So(book.ID, ShouldEqual, id)
		So(book.Price, ShouldEqual, price)
	})
}

func TestSetPrice(t *testing.T) {
	testCases := []struct {
		desc  string
		id    int
		price float64
	}{
		{
			desc:  "正数价格",
			id:    1,
			price: 25.6,
		},
		{
			desc:  "负数价格",
			id:    2,
			price: -11,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if book := new(Book).Builder(tC.id).SetPrice(tC.price).Build(); book.Price < 0 {
				t.Errorf("SetPrice() = %v, want %v", tC.price, 0)
			}
		})
	}
}
