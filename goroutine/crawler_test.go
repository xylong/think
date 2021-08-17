package goroutine

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var (
	url  = "https://news.cnblogs.com/n/page/%d/"
	page = 3
)

func TestCrawler_Get(t *testing.T) {
	Convey("网页爬取", t, func() {
		c := newCrawler(url)
		ShouldBeNil(c.Get(page))
	})
}
