package main

import "fmt"

func main() {
	// 使用 Wire 自动生成的依赖注入器
	userService := InitializeUserService()

	// 调用 UserService 的方法获取用户名字
	fmt.Println(userService.GetUserName(1))
}
