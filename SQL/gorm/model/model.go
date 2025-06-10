package model

import (
	"time"
)

/*
CREATE TABLE `student` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(16) DEFAULT NULL,
  `age` tinyint(3) unsigned DEFAULT NULL,
  `gender` tinyint(1) DEFAULT NULL,
  `birthday` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
*/

type Student struct {
	ID       int64     `gorm:"primaryKey"`
	Name     string    `gorm:"column:name"`
	Age      int       `gorm:"column:age"`
	Gender   int       `gorm:"column:gender"`
	Email    *string   `gorm:"column:email"`
	Birthday time.Time `gorm:"column:birthday"`
}

func (*Student) TableName() string {
	return "student"
}

type Menu struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	ParentID *int
	Children []*Menu `gorm:"foreignKey:ParentID;references:ID"`
}
