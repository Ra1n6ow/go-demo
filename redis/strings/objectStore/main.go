package main

/*
结构体存储:
	手动Marshal结构体
*/

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type UserName struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	Age      int64  `json:"age"`
}

func (u *UserName) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u) //解析
}

func (u *UserName) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u) //反解析
}

func main() {
	rdb := RedisConnect()

	data := &UserName{
		UserId:   1,
		UserName: "LiuBei",
		Age:      28,
	}

	/*
		//手动Marshal结构体

		//结构体转为json字串的[]byte
		b, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
		}

		//写入
		err = rdb.Set(ctx, "struct1", string(b), time.Minute*2).Err()
		if err != nil {
			fmt.Println("err: ", err)
		}

		//查找
		result := rdb.Get(ctx, "struct1")
		fmt.Println(result.Val())
	*/

	//写入,如果 data 对象实现了 MarshalBinary，在写入前会自动调用该方法
	err := rdb.Set(ctx, "struct1", data, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	//读出，如果 result 对象实现了 UnmarshalBinary，则在读取后会自动调用该方法
	result := &UserName{}
	err = rdb.Get(ctx, "struct1").Scan(result)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("get success: %+v\n", result)
	fmt.Println(data)
}

func RedisConnect() (rdb *redis.Client) {
	var myRedis redis.Options
	myRedis.Addr = "127.0.0.1:6379"
	myRedis.Password = ""
	//myRedis.DB = 1
	rdb = redis.NewClient(&myRedis)
	return rdb
}
