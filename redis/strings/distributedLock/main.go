package main

/*
使用 redis 的 NX 实现分布式锁
多实例见：https://www.jb51.net/article/235277.htm
*/

import (
	"context"
	"errors"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

var client *redis.Client

// 当key获取的值等于传入的value时，代表是自己持有，然后执行删除
const unlockScript = `
if redis.call("get",KEYS[1]) == ARGV[1] then
  return redis.call("del",KEYS[1])
else
  return 0
end`

func lottery(ctx context.Context) error {
	// 加锁
	myRandomValue := gofakeit.UUID()
	resourceName := "resource_name"
	ok, err := client.SetNX(ctx, resourceName, myRandomValue, time.Second*30).Result()
	if err != nil {
		return err
	}

	// 如果SetNX 失败，说明key已经存在
	if !ok {
		return errors.New("系统繁忙，请重试")
	}

	// 解锁
	defer func() {
		script := redis.NewScript(unlockScript)
		script.Run(ctx, client, []string{resourceName}, myRandomValue)
	}()

	// 业务处理
	time.Sleep(time.Second)
	return nil
}

func main() {
	client = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
		// context deadline exceeded
		// time.Sleep(time.Second * 3)
		err := lottery(ctx)
		if err != nil {
			fmt.Println("111")
			fmt.Println(err)
		}
	}()
	go func() {
		defer wg.Done()
		ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
		// context deadline exceeded
		// time.Sleep(time.Second * 3)
		err := lottery(ctx)
		if err != nil {
			fmt.Println("222")
			fmt.Println(err)
		}
	}()
	wg.Wait()
}
