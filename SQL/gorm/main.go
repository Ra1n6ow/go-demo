package main

import (
	"fmt"

	"github.com/ra1n6ow/go-demo/SQL/gorm/relations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
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
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	// db.AutoMigrate(&model.Student{})
	// db.AutoMigrate(&model.Menu{})

	// One
	// one.InsertOne(db)
	// one.InsertMany(db)
	// one.QueryOne(db)
	// one.UpdateOne(db)
	// one.DeleteOne(db)
	// advanced.InitData(db)
	// advanced.Query(db)
	// advanced.Select(db)
	// advanced.Other(db)
	relations.CreateChildren(db)
	// relations.QueryChildren(db)
	// relations.QueryParent(db)
}
