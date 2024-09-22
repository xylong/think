package lib

import "fmt"

type Locker struct {
	Key string
}

func NewLocker(key string) *Locker {
	return &Locker{Key: key}
}

// Lock 上锁
func (l *Locker) Lock() *Locker {
	cmd := redisClient.SetNX(l.Key, 1, 0)
	if ok, err := cmd.Result(); err != nil || !ok {
		panic(fmt.Sprintf("lock error with key [%s]", l.Key))
	}

	return l
}

// Unlock 解锁
func (l *Locker) Unlock() *Locker {
	redisClient.Del(l.Key)

	return l
}
