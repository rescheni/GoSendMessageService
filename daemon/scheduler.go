// scheduler.go
package daemon

import (
	basic "GoMessageService/Basic"
	"GoMessageService/plug-in/fate"
	"GoMessageService/sendserver"
	"strings"
)

// 定义守护进程函数，定时任务执行
func StartDailyScheduler() {
	// 加载配置
	basic.LoadConfig()

	// // 创建一个每分钟触发的 ticker
	// ticker := time.NewTicker(1 * time.Minute)
	// defer ticker.Stop()

	// for {
	// 	<-ticker.C
	// 	// 获取当前时间
	// 	now := time.Now()

	// 	// 如果当前时间是 8 点整
	// 	if now.Hour() == 9 && now.Minute() == 23 {
	// 		// 执行发送邮件的操作
	// 		emailContent := strings.Join(fate.FateTest(), "\r\n")
	// 		// sendserver.EmailSend([]string{"1413024330@qq.com"}, emailContent, "Fate")
	// 		sendserver.ServerJiang(emailContent, "Fate")
	// 		// 打印日志
	// 		fmt.Println("Email sent successfully at", now.Format("2006-01-02 15:04:05"))
	// 	} else {
	// 		// 不是 8 点，继续等待
	// 		fmt.Println("Waiting for 8:00 AM...")
	// 	}
	// }

	emailContent := strings.Join(fate.FateTest(), "\r\n")
	// sendserver.EmailSend([]string{"1413024330@qq.com"}, emailContent, "Fate")
	// sendserver.ServerJiang("Fate", emailContent)
	// sendserver.ServerFeiShu("Fate", emailContent)
	sendserver.DingServer(emailContent)
	// sendserver.WxPusherCreateQRCode()
	// sendserver.SendWxPusher("Fate", emailContent)
	// sendserver.WxPusherUserList()

}
