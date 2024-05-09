package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
)

var ctx = context.Background()

const CART_PREFIX = "cart_"

func add(client *redis.Client, userId int, productId string) {
	res := client.HSet(ctx, fmt.Sprintf("%s%d", CART_PREFIX, userId), productId, 1)
	if res.Err() != nil {
		panic(res.Err())
	}

	if res.Val() == 1 {
		fmt.Println("添加购物车成功,ProductID: ", productId)
	}
}

func del(client *redis.Client, userId int, productId string) {
	res := client.HDel(ctx, fmt.Sprintf("%s%d", CART_PREFIX, userId), productId)
	if res.Val() == 1 {
		fmt.Println("商品删除成功,ProductID: ", productId)
	}
}

func incr(client *redis.Client, userId int, productId string) {
	res, _ := client.HIncrBy(ctx, fmt.Sprintf("%s%d", CART_PREFIX, userId), productId, 1).Result()
	fmt.Println(productId, "商品数量加1成功，剩余数量为: ", res)
}

func decr(client *redis.Client, userId int, productId string) {
	res, _ := client.HGet(ctx, fmt.Sprintf("%s%d", CART_PREFIX, userId), productId).Result()
	value, err := strconv.Atoi(res)
	if err != nil {
		fmt.Println("无法将 res 转换为整数类型:", err)
		panic(err)
	}
	if value-1 <= 0 {
		// 购物车余额为0，删除商品
		del(client, userId, productId)
		fmt.Println("商品 productID: ", productId, " 为0，删除...")
		return
	}
	deRes, _ := client.HIncrBy(ctx, fmt.Sprintf("%s%d", CART_PREFIX, userId), productId, -1).Result()
	fmt.Println("商品数量减1成功，剩余数量为: ", deRes)
}

func submitOrder(client *redis.Client, userId int) {
	res := client.HGetAll(ctx, fmt.Sprintf("%s%d", CART_PREFIX, userId))
	fmt.Println("用户: ", userId, " 提交订单")
	// 检查错误
	if err := res.Err(); err != nil {
		panic(err)
	}

	// 获取结果
	data, err := res.Result()
	if err != nil {
		panic(err)
	}

	// 打印结果
	for field, value := range data {
		fmt.Printf("%s: %s\n", field, value)
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

	defer client.Close()

	// 模拟操作
	add(client, 1024, "83694")
	add(client, 1024, "1273979")
	add(client, 1024, "123323")
	submitOrder(client, 1024)
	del(client, 1024, "123323")

	// 增减购物车
	incr(client, 1024, "83694")
	decr(client, 1024, "1273979")
}
