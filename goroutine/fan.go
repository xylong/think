package goroutine

import "sync"

// FanIn 扇入
// *多个channel(相同类型)读取数据，合并输出到一个总的channel里，然后读取出来作进一步操作（对顺序没有要求），直到这些通道关闭
func FanIn(cs ...<-chan interface{}) <-chan interface{} {
	result := make(chan interface{})
	wg := sync.WaitGroup{}

	for _, c := range cs {
		wg.Add(1)
		go func(c <-chan interface{}) {
			defer wg.Done()
			for v := range c {
				result <- v
			}
		}(c)
	}

	// !因为不知道什么时候遍历结束，所以不知道什么时候结束，所以🉐得塞到协程等wait结束时来关闭
	go func() {
		defer close(result)
		wg.Wait()
	}()

	return result
}

// FanOut 扇出
func FanOut(data <-chan interface{}, worker ...chan interface{}) {
	go func() {
		defer func() {
			for _, c := range worker {
				close(c)
			}
		}()

		for v := range data {
			for _, c := range worker {
				c<-v
			}
		}
	}()
}
