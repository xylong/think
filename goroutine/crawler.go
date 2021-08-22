package goroutine

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// crawler 爬虫
type crawler struct {
	url string
}

func newCrawler(url string) *crawler {
	return &crawler{url: url}
}

// Get 爬取网页内容
func (c *crawler) Get(page int) {
	wg := sync.WaitGroup{}
	// 任务完成通知
	done := make(chan struct{})
	// 内容
	contentChan := make(chan map[int][]byte, page)

	// 协程爬取
	for i := 1; i <= page; i++ {
		wg.Add(1)
		go func(index int) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println(err)
				}
				wg.Done()
			}()
			url := fmt.Sprintf(c.url, index)
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}

			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			contentChan <- map[int][]byte{index: bytes}
		}(i)
	}

	// 在协程中关闭channel，防止主线程一直运行
	go func() {
		defer close(contentChan)
		wg.Wait()
		done <- struct{}{}
	}()

	// 写入文件
writeLoop:
	for {
		select {
		case content := <-contentChan:
			for index, item := range content {
				if err := ioutil.WriteFile(fmt.Sprintf("%d.html", index), item, 0644); err != nil {
					fmt.Println(err.Error())
				}
			}
		case <-time.After(time.Second * 2):
			break writeLoop
		case <-done:
			break writeLoop
		}
	}
}
