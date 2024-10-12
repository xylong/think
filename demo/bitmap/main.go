package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var (
	rdb *redis.Client
)

func init() {
	rdb = RedisClient()
	if rdb == nil {
		panic("redis client is nil")
	}
}

func RedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// 签到
func signIn(key string, dayOffset int64) bool {
	i, err := rdb.SetBit(context.Background(), key, dayOffset, 1).Result()
	// 签到失败
	if err != nil {
		return false
	}
	// 重复签到
	if i == 1 {
		return false
	}

	return true
}

// 获取签到
func getSign(key string, dayOffset int64) int64 {
	i, err := rdb.GetBit(context.Background(), key, dayOffset).Result()
	if err != nil {
		return 0
	}

	return i
}

// 统计当月的签到情况
// days 显示多少天的签到，offset 从第几天开始
// getSignOfMonth(1, 2024, 31, 0)
func getSignOfMonth(userID int, year int, days, offset int) ([]bool, error) {
	typ := fmt.Sprintf("u%d", days)
	key := fmt.Sprintf("user:%d:%d", year, userID)

	s, err := rdb.BitField(context.Background(), key, "GET", typ, offset).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get bitfield: %w", err)
	}

	if len(s) != 0 {
		signInBits := s[0]
		signInSlice := make([]bool, days)
		for i := 0; i < days; i++ {
			signInSlice[i] = (signInBits & (1 << (days - 1 - i))) != 0
		}
		return signInSlice, nil
	} else {
		return nil, errors.New("no result returned from BITFIELD command")
	}
}

// 获取指定年份的累计签到天数
func getCumulativeDays(userID, year int, dayOfYear int) (int, error) {
	var (
		key            = fmt.Sprintf("user:%d:%d", year, userID)
		segmentSize    = 63
		cumulativeDays = 0
		bitOps         = make([]any, 0)
	)

	for i := 0; i < dayOfYear; i += segmentSize {
		size := segmentSize
		if i+segmentSize > dayOfYear {
			size = dayOfYear - i
		}

		bitOps = append(bitOps, "GET", fmt.Sprintf("u%d", size), fmt.Sprintf("#%d", i))
	}

	values, err := rdb.BitField(context.Background(), key, bitOps...).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to get bitfield: %w", err)
	}

	for idx, value := range values {
		if value != 0 {
			size := segmentSize
			if (idx+1)*segmentSize > dayOfYear {
				size = dayOfYear % segmentSize
			}
			for j := 0; j < size; j++ {
				if (value & (1 << (size - 1 - j))) != 0 {
					cumulativeDays++
				}
			}
		}
	}

	return cumulativeDays, nil
}

func main() {
	//res := signIn("user:2024:1", 0) // 用户1在 2024-01-01 签到
	//fmt.Println(res)
	//res = signIn("user:2024:1", 0) // 重复签到
	//fmt.Println(res)

	//fmt.Println(getSign("user:2024:1", 0))

	fmt.Println(getSignOfMonth(1, 2024, 30, 0))
	fmt.Println(getCumulativeDays(1, 2024, 10))
}
