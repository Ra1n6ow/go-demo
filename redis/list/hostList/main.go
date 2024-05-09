package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"time"
)

/*
hot-list 的 List 来存储热点话题列表。初始化的时候，先写进去 10 条热点新闻，
然后启动一个 changeHotList 的协程，随机选择一条热点新闻，用 LSET 命令进行修改。
最后主线程会模拟客户端，每次取热点列表的时候，会直接用 LRANGE 命令把整个 hot-list 里面的热点新闻列表全部取出来进行展示。
*/
var ctx = context.Background()

// 随机更改一条热点新闻
func changeHotList(client *redis.Client) {
	for {
		rand.New(rand.NewSource(time.Now().UnixMilli()))
		randomNumber := rand.Intn(10)
		err := client.LSet(ctx, "host-list", int64(randomNumber),
			fmt.Sprintf("第 %d 条热点新闻，更新时间:%s", randomNumber, time.Now().Format("2006-01-02 15:04:05"))).Err()
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 2)
	}
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis服务器地址
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	// 初始化10条热点数据
	for i := 0; i < 10; i++ {
		err := client.RPush(ctx, "host-list",
			fmt.Sprintf("第 %d 条热点新闻，更新时间:%s", i, time.Now().Format("2006-01-02 15:04:05"))).Err()
		if err != nil {
			fmt.Println(err)
		}
	}

	go changeHotList(client)

	// 模拟客户端拉取热点新闻5次
	for i := 0; i < 5; i++ {
		res, err := client.LRange(ctx, "host-list", 0, -1).Result()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("热点新闻：", res)
		time.Sleep(time.Second * 5)
	}
}
