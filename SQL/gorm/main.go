package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/ra1n6ow/go-demo/SQL/gorm/advanced"
)

func main() {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s`,
		"root",
		"123456",
		"127.0.0.1:3306",
		"test",
		true,
		"UTC")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	// db.AutoMigrate(&model.Student{})

	// One
	// one.InsertOne(db)
	// one.InsertMany(db)
	// one.QueryOne(db)
	// one.UpdateOne(db)
	// one.DeleteOne(db)
	// advanced.InitData(db)
	// advanced.Query(db)
	// advanced.Select(db)
	advanced.Other(db)
}
