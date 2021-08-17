package goroutine

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// crawler 爬虫
type crawler struct {
	url string
}

func newCrawler(url string) *crawler {
	return &crawler{url: url}
}

// Get 爬取网页内容
func (c *crawler) Get(page int) error {
	errChan := make(chan error)
	contentChan := make(chan map[int][]byte)
	wg := sync.WaitGroup{}

	for i := 1; i <= page; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			url := fmt.Sprintf(c.url, index)
			resp, err := http.Get(url)
			if err != nil {
				errChan <- err
			}

			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				errChan <- err
			}
			contentChan <- map[int][]byte{index: bytes}
		}(i)
	}

	// 在协程中关闭channel，防止主线程一直运行
	go func() {
		defer close(contentChan)
		defer close(errChan)
		wg.Wait()
	}()

	for item := range contentChan {
		for i, v := range item {
			if err := ioutil.WriteFile(fmt.Sprintf("%d.html", i), v, 0644); err != nil {
				errChan <- err
			}
		}
	}

	return <-errChan
}
