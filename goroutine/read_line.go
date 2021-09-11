package goroutine

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

// reader 逐行读取
type reader struct {
}

func newReader() *reader {
	return &reader{}
}

func (r *reader) Read() {
	file, err := os.OpenFile("./README.md", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	fw := bufio.NewReader(file)
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			for {
				lock.Lock()
				str, err := fw.ReadString('\n')
				if err != nil {
					if err == io.EOF {
						lock.Unlock()
						break
					}
					log.Fatalln(err)
				}
				time.Sleep(time.Millisecond * 200)
				fmt.Printf("[协程%d]%s", index, str)
				lock.Unlock()
			}
		}(i)
	}

	wg.Wait()
}
