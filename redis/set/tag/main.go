package main

/*
准备几款裙子作为商品库，然后给每个商品打上标签，例如，商品 ID 是 1 的连衣裙， 是夏天穿的短款。
给商品打好标签之后，我们开始给用户打标签，线上用户的标签一般是需要结合用户历史浏览数据、用户订单等一系列信息推算出来的。
例如，一个用户自身的标签是“女性”“20~30 岁之间”“爱网购”，然后会最近经常搜“裙子”“优惠券”这些关键字，而现在也是夏天，
那系统可能就认为小伙伴有极大可能会要买裙子。直接给用户 1 打上“夏款”“折扣”的标签，
然后算一下用户 1 和各个商品的标签交集，有交集的话，就给这个用户推荐这个商品了。
*/

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
)

var ctx = context.Background()

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis服务器地址
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	defer client.Close()

	productDB := map[int]string{
		1: "XX连衣裙",
		2: "XX长裙",
		3: "XX半身裙",
	}

	// tag_p前缀加商品ID作为key
	client.SAdd(ctx, "tag_p_1", "短款", "A字裙", "夏款")
	client.SAdd(ctx, "tag_p_2", "春款", "长袖")
	client.SAdd(ctx, "tag_p_3", "过膝款", "纯棉", "折扣")

	// tag_u前缀加用户ID作为Key
	client.SAdd(ctx, "tag_u_1", "夏款", "折扣")

	for productID, productName := range productDB {
		res, _ := client.SInter(ctx, "tag_p_"+strconv.Itoa(productID), "tag_u_1").Result()
		if len(res) != 0 {
			fmt.Println("精选页推荐: ", productName, ", 推荐原因：", res)
		}
	}
}
