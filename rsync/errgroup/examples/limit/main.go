package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	// 创建一个 errgroup.Group
	var eg errgroup.Group
	// 设置最大并发限制为 3
	eg.SetLimit(3)

	// 启动 10 个 goroutine
	for i := 1; i <= 10; i++ {
		eg.Go(func() error {
			// 打印正在运行的 goroutine
			fmt.Printf("Goroutine %d is starting\n", i)
			time.Sleep(2 * time.Second) // 模拟任务耗时
			fmt.Printf("Goroutine %d is done\n", i)
			return nil
		})
	}

	// 等待所有 goroutine 完成
	if err := eg.Wait(); err != nil {
		fmt.Printf("Encountered an error: %v\n", err)
	}

	fmt.Println("All goroutines complete.")
}
