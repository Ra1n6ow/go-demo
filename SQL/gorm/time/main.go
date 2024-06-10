package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

/*
	生成的 sql

CREATE TABLE `test_times` (

	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	`tested_at` datetime(3) DEFAULT NULL,
	`created_at` timestamp NOT NULL DEFAULT current_timestamp(),
	`update_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
	PRIMARY KEY (`id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
*/
type TestTime struct {
	ID        uint      `gorm:"primaryKey"`
	TestedAt  time.Time `gorm:"column:tested_at"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP;<-:create" json:"created_at,omitempty"`
	UpdateAt  time.Time `gorm:"column:update_at;type:timestamp;default:current_timestamp  on update current_timestamp"`
}

func main() {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s`,
		"root",
		"123456",
		"127.0.0.1:3306",
		"test",
		true,
		"UTC")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 1. Auto migration for given models
	db.AutoMigrate(&TestTime{})
}
