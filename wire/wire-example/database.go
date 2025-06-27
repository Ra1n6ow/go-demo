package main

import "fmt"

// Database 模拟数据库连接
type Database struct {
	DSN string // 数据源名称
}

// NewDatabase 是 Database 的构造函数
func NewDatabase() *Database {
	fmt.Println("Initializing database connection...")
	return &Database{
		DSN: "user:password@tcp(127.0.0.1:3306)/example_db",
	}
}
