package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dg *gorm.DB

// InitDB 初始化数据库
func InitDB() {
	var err error
	dg, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect database", err)
	}

	// Migrate the schema
	// 目前没有打算注册账号，所以先注释掉
	// dg.AutoMigrate(&User{})
	dg.AutoMigrate(&Cron{})
}
