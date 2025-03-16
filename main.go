package main

import (
	api "GoMessageService/API"
	"fmt"
)

//	func main() {
//		basic.LoadConfig()
//		sendserver.EmailSend([]string{"1413024330@qq.com"}, strings.Join(fate.FateTest(), "\r\n"), "Fate")
//	}

func main() {
	// 启动每天的定时任务守护进程
	fmt.Println("Starting the Task...")
	// 开启监听api
	api.APIStart()

	// go daemon.StartDailyScheduler()

	// // 启动定时任务
	// basic.SetCronTask("0 */1 * * * ?", func() {
	// })

	// sendserver.Send_group_msg(yiyanAPI.GetSentence(11), "720057024")
	// 防止主线程退出
	select {}
}
