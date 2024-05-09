package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()

func rateLimiter(rdb *redis.Client, key string, limit int, seconds int) bool {
	// 使用Redis的INCR命令递增键的值
	val, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	// 如果键不存在，设置过期时间，并将计数器初始化为1
	if val == 1 {
		_, err = rdb.Expire(ctx, key, time.Duration(seconds)*time.Second).Result()
		if err != nil {
			panic(err)
		}
	}

	// 如果计数器的值超过了限制，则返回false表示超出限制
	if int(val) > limit {
		return false
	}
	return true
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// 测试限流
	key := "my_rate_limiter"
	limit := 5
	seconds := 10

	for i := 0; i < 10; i++ {
		if rateLimiter(rdb, key, limit, seconds) {
			fmt.Printf("Request %d is allowed.\n", i)
		} else {
			fmt.Printf("Request %d is denied.\n", i)
		}
	}
}
