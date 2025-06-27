package main

import "fmt"

// UserService 表示用户业务逻辑服务
type UserService struct {
	repo *UserRepository
}

// NewUserService 是构造函数，实例化 UserService 并依赖 UserRepository
func NewUserService(repo *UserRepository) *UserService {
	fmt.Println("Initializing UserService...")
	return &UserService{repo: repo}
}

// GetUserName 通过 UserService 获取用户的名字
func (s *UserService) GetUserName(id int) string {
	return s.repo.GetNameByID(id)
}
