//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// InitializeUserService 是 Wire 的注入器函数
// 它会解析依赖图并生成依赖初始化代码
func InitializeUserService() *UserService {
	wire.Build(
		NewDatabase,       // Database 的提供者
		NewUserRepository, // UserRepository 的提供者
		NewUserService,    // UserService 的提供者
	)
	return nil
}
