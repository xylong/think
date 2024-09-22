package lib

import (
	"fmt"
	"time"
)

type Locker struct {
	key    string
	expire time.Duration
}

func NewLocker(key string) *Locker {
	return &Locker{key: key}
}

func NewLockerWithTTL(key string, expire time.Duration) *Locker {
	return &Locker{
		key:    key,
		expire: expire,
	}
}

// Lock 上锁
func (l *Locker) Lock() *Locker {
	cmd := redisClient.SetNX(l.key, 1, l.expire)
	if ok, err := cmd.Result(); err != nil || !ok {
		panic(fmt.Sprintf("lock error with key [%s]", l.key))
	}

	return l
}

// Unlock 解锁
func (l *Locker) Unlock() *Locker {
	redisClient.Del(l.key)

	return l
}
