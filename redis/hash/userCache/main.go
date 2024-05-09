package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func newUser(name, password, phone string) map[string]string {
	data := map[string]string{
		"name":     name,
		"password": password,
		"phone":    phone,
	}
	return data
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis服务器地址
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	defer client.Close()

	// 手动生成点数据
	//client.HSet(ctx, "13188888888", newUser("dagou", "123456", "13188888888"))
	//client.HSet(ctx, "13288888888", newUser("ergou", "123456", "13288888888"))
	//client.HSet(ctx, "13388888888", newUser("sangou", "123456", "13388888888"))

	res, err := client.HGet(ctx, "13388888888", "name").Result()
	//res, err := client.HGet(ctx, "13488888888", "phone").Result()
	if err == redis.Nil {
		// 未命中处理
		fmt.Println("未命中")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}
}
