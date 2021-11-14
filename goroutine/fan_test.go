package goroutine

import (
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/smartystreets/goconvey/convey"
)

func TestFanIn(t *testing.T) {
	convey.Convey("扇入模式", t, func() {
		r := gin.New()

		r.GET("user", func(c *gin.Context) {
			// ? 分块输出
			c.Writer.Header().Add("Transfer-Encoding", "chunked")
			c.Writer.WriteHeader(http.StatusOK)

			ch := FanIn(getUserInfo(), getUserBalance())
			for v := range ch {
				c.Writer.Write([]byte(v.(string)))
				c.Writer.(http.Flusher).Flush() // ? 刷新缓冲区
			}
		})

		r.Run()
	})
}

func getUserInfo() <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer close(ch)
		time.Sleep(time.Second * 2) // 模拟耗时任务
		ch <- "用户🆔：" + time.Now().Format("150405")
	}()

	return ch
}

func getUserBalance() <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer close(ch)
		time.Sleep(time.Second * 5) // 模拟耗时任务
		ch <- "用户余额：" + time.Now().Format("150405")
	}()

	return ch
}
