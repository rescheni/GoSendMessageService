package main

import (
	api "GoMessageService/API"
	"GoMessageService/database"
	log "GoMessageService/log"
)

// send all the messages
// func sendall() {
// 	// 发送所有消息
// 	// 加载配置
// 	basic.LoadConfig()

// 	title := "GoMessageService Start"
// 	emailContent := "这是一封测试邮件，GoMessageService已经启动 \n"
// 	// 发送消息
// 	sendserver.SendEmail([]string{"1413024330@qq.com"}, title, emailContent)
// 	sendserver.SendServerJiang(title, emailContent)
// 	sendserver.SendFeiShu(title, emailContent)
// 	sendserver.SendDing(title, emailContent)
// 	sendserver.WxPusherCreateQRCode()
// 	sendserver.SendWxPusher(title, emailContent)
// 	sendserver.WxPusherUserList()

// }

func main() {
	// 启动每天的定时任务守护进程
	log.Logger.Info("GoMessageService Start")

	// 初始化数据库
	database.InitDB()
	// 加载定时任务
	database.LoadCornTaskOnDb()

	// 开启监听api服务 启动gin
	api.APIStart()
	// 防止主线程退出
	select {}
}
