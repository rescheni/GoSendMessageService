package database

import (
	log "GoMessageService/log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dg *gorm.DB

func GetDB() *gorm.DB {
	return dg
}

func InsertCron(cron *Cron) {
	dg.Create(cron)
}

func UpdateCron(cron *Cron) {
	dg.Save(cron)
}

func DeleteCron(entryID string) {
	var cron Cron
	dg.Where("entry_id = ?", entryID).First(&cron)
	dg.Delete(cron)
}

func GetCronList() []Cron {
	var crons []Cron
	dg.Find(&crons)

	log.Logger.Info("GetCronList", crons)

	return crons
}

func GetCronByID(id int) Cron {
	var cron Cron
	dg.First(&cron, id)
	return cron
}

// InitDB 初始化数据库
func InitDB() {
	var err error
	dg, err = gorm.Open(sqlite.Open("cron.db"), &gorm.Config{})
	if err != nil {
		log.Logger.Fatal("Failed to connect database", err)
	}

	// Migrate the schema
	// 目前没有打算注册账号，所以先注释掉
	// dg.AutoMigrate(&User{})
	dg.AutoMigrate(&Cron{})
}

// CloseDB 关闭数据库
func CloseDB() {
	db, err := dg.DB()
	if err != nil {
		log.Logger.Fatal("Failed to get database", err)
	}
	db.Close()
}
