package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

/*
模拟点赞功能
*/
var ctx = context.Background()
var videoId = 1089
var listName = fmt.Sprintf("like-list-%d", videoId)

// 点赞
func createLike(client *redis.Client) {
	for i := 0; i < 10000; i++ {
		err := client.RPush(ctx, listName, i).Err()
		if err != nil {
			panic(err)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func readLike(client *redis.Client) {
	for {
		results, err := client.LRange(ctx, listName, 0, 1000).Result()
		if err != nil {
			panic(err)
		}
		time.Sleep(2 * time.Second)
		fmt.Printf("小伙伴 %s 点赞了视频 %d \n", results, videoId)
		rmRes, err := client.LTrim(ctx, listName, int64(len(results)), -1).Result()
		if err != nil {
			panic(err)
		}
		if rmRes == "OK" {
			fmt.Println("删除成功")
		}
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

	go createLike(client)
	go readLike(client)

	// 等 60 s
	time.Sleep(60 * time.Second)
}
