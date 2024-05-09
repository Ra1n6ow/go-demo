package main

/*
使用主线模拟一个用户的购物操作，这个用户每隔一秒调用一下 addProduct 方法。
addProduct() 里面做两件事：
一件事是检查用户的 userId 是否在 blackUserIds 这个黑名单中，如果存在的话，就会拦截用户的购物操作；
另一件事是在黑名单检查通过之后，把商品 id 添加到当前用户的购物车中。

在黑名单检查没通过的时候，这里做了一个简化处理，就是抛出异常，之后会调用 faceCheck() 方法，模拟用户做人脸验证的操作。
在人脸识别成功之后，就会把当前的userId 从 blackUserIds 黑名单里面删掉，这个用户就可以继续正常购物了。
*/

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

var ctx = context.Background()

func addProduct(client *redis.Client, userId, productId int) error {
	isBlack, _ := client.SIsMember(ctx, "blackUserIds", strconv.Itoa(userId)).Result()
	if isBlack {
		fmt.Println(userId, "帐号存在风险，请先去完成人脸验证...")
		return errors.New("黑名单用户")
	}
	res, _ := client.HSet(ctx, fmt.Sprintf("cart_%d", userId), strconv.Itoa(productId), "1").Result()
	if res != 0 {
		fmt.Println("添加购物车成功, ProductId: ", productId)
	}
	return nil
}

func faceCheck(client *redis.Client, userId int) {
	fmt.Println("人脸验证中。。。")
	time.Sleep(time.Second * 10)
	res, _ := client.SRem(ctx, "blackUserIds", strconv.Itoa(userId)).Result()
	if res == 1 {
		fmt.Println("人脸验证完成。。。")
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

	userId := 5

	// 模拟5秒后，将 5 号用户加入黑名单
	go func() {
		time.Sleep(time.Second * 5)
		client.SAdd(ctx, "blackUserIds", strconv.Itoa(userId))
	}()

	for i := 0; i < 10; i++ {
		err := addProduct(client, userId, i)
		if err != nil {
			faceCheck(client, userId)
		}
		time.Sleep(time.Second * 1)
	}
}
