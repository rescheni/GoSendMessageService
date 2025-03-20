package database

import (
	basic "GoMessageService/Basic"
	log "GoMessageService/log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dg *gorm.DB //全局的数据库实例

// 获取数据库实例
func GetDB() *gorm.DB {
	return dg
}
func SaveSendedMessage(sendtime string, title string, message string, sendType string, userid string) {
	var sendMessage SendMessage
	sendMessage.SendTime = sendtime
	sendMessage.SendType = sendType
	sendMessage.UserId = userid
	sendMessage.Title = title
	sendMessage.Message = message
	dg.Create(&sendMessage)

}
func GetSendedMessage() []SendMessage {
	var sendMessages []SendMessage
	dg.Find(&sendMessages)
	return sendMessages
}
func DeleteSendedMessage(sendid int) {

	var sendMessage SendMessage
	dg.First(&sendMessage, sendid)
	dg.Delete(sendMessage)
}

// 插入cron定时任务
func InsertCron(cron *Cron) bool {
	// TODO: 检查cron的有效性
	// 检查 entryID 是否为空
	if cron.EntryID == "" {
		log.Logger.Error("EntryID is empty")
		return false
	}

	// 检查 cron ID 是否重复
	var existingCron Cron
	result := dg.Where("entry_id = ?", cron.EntryID).First(&existingCron)
	if result.RowsAffected > 0 {
		log.Logger.Error("Cron ID already exists", cron.EntryID)
		return false
	}

	dg.Create(cron)
	return true
}

// 更新cron定时任务  目前没有使用到
func UpdateCron(cron *Cron) {
	dg.Save(cron)
}

// 删除cron定时任务
func DeleteCron(entryID string) {
	var cron Cron
	dg.Where("entry_id = ?", entryID).First(&cron)
	dg.Delete(cron)
}

// 获取cron定时任务列表
func GetCronList() []Cron {
	var crons []Cron
	dg.Find(&crons)

	log.Logger.Info("GetCronList", crons)

	return crons
}

// 通过id 获取cron定时任务
func GetCronByID(id int) Cron {
	var cron Cron
	dg.First(&cron, id)
	return cron
}

// InitDB 初始化数据库
func InitDB() {
	cfg := basic.LoadConfig()
	var err error
	dg, err = gorm.Open(sqlite.Open(cfg.Sqlite.Db_path), &gorm.Config{})
	if err != nil {
		log.Logger.Fatal("Failed to connect database", err)
	}

	// Migrate the schema
	// 目前没有打算注册账号，所以先注释掉
	// dg.AutoMigrate(&User{})
	dg.AutoMigrate(&Cron{})
	dg.AutoMigrate(&SendMessage{})
}

// CloseDB 关闭数据库
func CloseDB() {
	db, err := dg.DB()
	if err != nil {
		log.Logger.Fatal("Failed to get database", err)
	}
	db.Close()
}
