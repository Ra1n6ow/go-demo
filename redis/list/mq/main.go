package main

/*
模拟消息队列
*/

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()

func SendMessage(client *redis.Client, message string) error {
	err := client.LPush(ctx, "testlist", message).Err()
	if err != nil {
		return err
	}
	return nil
}

func ProcessMessages(client *redis.Client) {
	for {
		result, err := client.BRPop(ctx, time.Second, "testlist").Result()
		if err != nil {
			panic(err)
		}

		// 处理消息
		message := result[1]
		fmt.Println("Received message:", message)
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

	go ProcessMessages(client)

	// 发送消息
	err = SendMessage(client, "Hello, Redis!")
	if err != nil {
		panic(err)
	}

	// 等待消息处理完成
	time.Sleep(1 * time.Second)
}
