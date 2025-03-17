package main

import (
	api "GoMessageService/API"
	basic "GoMessageService/Basic"
	log "GoMessageService/log"
	"GoMessageService/sendserver"
)

// send all the messages
func sendall() {
	// 发送所有消息
	// 加载配置
	basic.LoadConfig()

	title := "GoMessageService Start"
	emailContent := "这是一封测试邮件，GoMessageService已经启动 \n"
	// 发送消息
	sendserver.EmailSend([]string{"1413024330@qq.com"}, emailContent, title)
	sendserver.ServerJiang(title, emailContent)
	sendserver.FeiShuSend(title, emailContent)
	sendserver.DingSend(emailContent)
	sendserver.WxPusherCreateQRCode()
	sendserver.SendWxPusher(title, emailContent)
	sendserver.WxPusherUserList()

}

func main() {
	// 启动每天的定时任务守护进程
	log.Logger.Info("GoMessageService Start")

	// 初始化数据库
	// TODO
	// 服务启动通知
	// sendall()

	// 开启监听api服务 启动gin
	api.APIStart()
	// 防止主线程退出
	select {}
}
