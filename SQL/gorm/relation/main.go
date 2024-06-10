package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

/*
不修改 join 表外键：
CREATE TABLE `user_languages` (
  `language_m_id` bigint(20) unsigned NOT NULL,
  `user_m_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`language_m_id`,`user_m_id`),
  KEY `fk_user_languages_user_m` (`user_m_id`),
  CONSTRAINT `fk_user_languages_language_m` FOREIGN KEY (`language_m_id`) REFERENCES `language` (`id`),
  CONSTRAINT `fk_user_languages_user_m` FOREIGN KEY (`user_m_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

修改后：
CREATE TABLE `user_languages` (
  `user_id` bigint(20) unsigned NOT NULL,
  `language_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`language_id`),
  KEY `fk_user_languages_language_m` (`language_id`),
  CONSTRAINT `fk_user_languages_language_m` FOREIGN KEY (`language_id`) REFERENCES `language` (`id`),
  CONSTRAINT `fk_user_languages_user_m` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
*/

type UserM struct {
	ID uint `gorm:"primaryKey"`
	//Languages []*LanguageM `gorm:"many2many:user_languages;"`
	Languages []*LanguageM `gorm:"many2many:user_languages;joinForeignKey:user_id;joinReferences:language_id"`
}

func (p *UserM) TableName() string {
	return "user"
}

type LanguageM struct {
	ID    uint     `gorm:"primaryKey"`
	Users []*UserM `gorm:"many2many:user_languages;joinForeignKey:language_id;joinReferences:user_id"`
}

func (p *LanguageM) TableName() string {
	return "language"
}

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
			//NoLowerCase: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 1. Auto migration for given models
	db.AutoMigrate(&UserM{}, &LanguageM{})
}
