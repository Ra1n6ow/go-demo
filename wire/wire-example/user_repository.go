package main

import "fmt"

// UserRepository 操作用户数据
type UserRepository struct {
	db *Database
}

// NewUserRepository 是构造函数，实例化 UserRepository 并依赖 Database
func NewUserRepository(db *Database) *UserRepository {
	fmt.Println("Initializing UserRepository...")
	return &UserRepository{db: db}
}

// GetNameByID 模拟通过用户 ID 获取用户名
func (r *UserRepository) GetNameByID(id int) string {
	return fmt.Sprintf("User #%d from database: %s", id, r.db.DSN)
}