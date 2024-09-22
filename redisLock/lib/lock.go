package lib

import "fmt"

// Lock 锁
func Lock(key string) {
	cmd := redisClient.SetNX(key, 1, 0)
	if ok, err := cmd.Result(); err != nil || !ok {
		panic(fmt.Sprintf("lock error with key [%s]", key))
	}
}
