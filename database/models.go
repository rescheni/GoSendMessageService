package database

import (
	basic "GoMessageService/Basic"
	log "GoMessageService/log"
	"GoMessageService/sendserver"
)

// type User struct {
// 	ID       uint   `json:"id" gorm:"primaryKey"`
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// 	Email    string `json:"email"`
// }

// cuncron 定时任务结构
// 数据持久化 cron 表
type Cron struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	ApiKey   string `json:"api" binding:"required"`
	CronExpr string `json:"cron_expr" binding:"required"`
	CronName string `json:"cron_name" binding:"required"`
	EntryID  string `json:"entry_id,omitempty"`
	Message  string `json:"message" binding:"required"`
	Title    string `json:"title,omitempty"`
	ToUser   string `json:"to_user,omitempty"`
	TaskType string `json:"task_type" binding:"required"` // 任务类型：wxpusher, dingding, server_jiang, email, feishu, napcat_qq
	// 任务状态：0-未启动，1-已启动
	// Status bool `json:"status"`
}

type SendMessage struct {
	// 自增id
	ID       uint   `json:"id" gorm:"primaryKey"`
	SendTime string `json:"send_time"`
	SendType string `json:"send_type"`
	Title    string `json:"title"`
	Message  string `json:"message"`
	UserId   string `json:"user"`
}

func LoadCornTaskOnDb() {

	log.Logger.Info("加载定时任务")

	crons := GetCronList()
	for _, cron := range crons {

		// 设置定时任务
		basic.SetCronTask(cron.CronName, cron.CronExpr, func() {
			switch cron.TaskType {
			case "wxpusher":
				sendserver.SendWxPusher(cron.Title, cron.Message)
			case "dingding":
				sendserver.SendDing(cron.Title, cron.Message)
			case "server_jiang":
				sendserver.SendServerJiang(cron.Title, cron.Message)
			case "email":
				sendserver.SendEmail([]string{cron.ToUser}, cron.Title, cron.Message)
			case "feishu":
				sendserver.SendFeiShu(cron.Title, cron.Message)
			case "napcat_qq":
				sendserver.SendQQPrivateMsg(cron.Message, cron.ToUser)
			}
		})
	}
}
