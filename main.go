package main

import (
	"GoMessageService/plug-in/yiyanAPI.go"
	"GoMessageService/sendserver"
	"fmt"
)

//	func main() {
//		basic.LoadConfig()
//		sendserver.EmailSend([]string{"1413024330@qq.com"}, strings.Join(fate.FateTest(), "\r\n"), "Fate")
//	}
func main() {
	// 启动每天的定时任务守护进程
	fmt.Println("Starting the scheduler...")
	// go daemon.StartDailyScheduler()

	sendserver.SendQQUserMessage(yiyanAPI.GetSentence(11), "1413024330")
	// 防止主线程退出
	// select {}
}
