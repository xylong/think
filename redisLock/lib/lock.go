package lib

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

const incrLua = `
if redis.call('get', KEYS[1]) == ARGV[1] then
  return redis.call('expire', KEYS[1],ARGV[2]) 				
 else
   return '0' 					
end`

type Locker struct {
	key        string
	expire     time.Duration
	unlock     bool
	incrScript *redis.Script
}

func NewLocker(key string) *Locker {
	return &Locker{
		key:        key,
		expire:     time.Second * 30,
		incrScript: redis.NewScript(incrLua),
	}
}

func NewLockerWithTTL(key string, expire time.Duration) *Locker {
	if expire <= 0 {
		panic("lock expire error")
	}

	return &Locker{
		key:        key,
		expire:     expire,
		incrScript: redis.NewScript(incrLua),
	}
}

// Lock 上锁
func (l *Locker) Lock() *Locker {
	cmd := redisClient.SetNX(l.key, 1, l.expire)
	if ok, err := cmd.Result(); err != nil || !ok {
		panic(fmt.Sprintf("lock error with key [%s]", l.key))
	}

	l.renewalLockTime()
	return l
}

// Unlock 解锁
func (l *Locker) Unlock() *Locker {
	if err := redisClient.Del(l.key).Err(); err != nil {
		panic(fmt.Sprintf("unlock error with key [%s]: %s", l.key, err.Error()))
	} else {
		l.unlock = true
	}

	return l
}

// 重置锁过期时间
func (l *Locker) resetExpire() {
	//redisClient.Expire(l.key, l.expire)
	result, err := l.incrScript.Run(redisClient, []string{l.key}, 1, l.expire.Seconds()).Result()
	log.Printf("key=%s ,续期结果:%v,%v\n", l.key, err, result)
}

// 锁过期时间续期
// 到锁过期2/3的时间才开始执行续期
func (l *Locker) renewalLockTime() {
	sleepTime := l.expire.Seconds() * 2 / 3

	go func() {
		for {
			time.Sleep(time.Duration(sleepTime) * time.Second)

			if l.unlock {
				break
			}

			l.resetExpire()
		}
	}()
}
