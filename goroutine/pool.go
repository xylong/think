package goroutine

import (
	"log"
	"sync"
)

type user struct {
	name string
}

func GetUser() *sync.Pool {
	return &sync.Pool{
		New: func() interface{} {
			log.Println("create user")
			return &user{
				name: "静静",
			}
		},
	}
}
