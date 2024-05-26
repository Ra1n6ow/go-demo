package main

import (
	"fmt"
	"time"
)

func main() {
	var u1 int64 = 720576000000
	t1 := time.UnixMilli(u1)
	fmt.Println("转换后的时间 t1: ", t1) // Println 显示本地的时区：东八区
	fmt.Println("转换后的时间 t1(UTC) : ", t1.UTC())
	f1 := t1.Format("2006-01-02")
	fmt.Println("转换后的时间字符串 f1: ", f1)
}
