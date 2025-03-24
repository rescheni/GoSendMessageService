// 项目api 接口文件
package api

import (
	basic "GoMessageService/Basic"
	"GoMessageService/services"

	"fmt"

	"github.com/gin-gonic/gin"
)

func APIStart() {
	cfg := basic.LoadConfig()
	apiConfig := cfg.Api

	router := gin.Default()

	send := router.Group("/send")
	// 发送消息
	{
		send.POST("/wxpusher", services.Send_wxpusher)
		send.POST("/dingding", services.Send_dingding)
		send.POST("/server jiang", services.Send_server_jiang)
		send.POST("/email", services.Send_email)
		send.POST("/feishu", services.Send_feishu)
		send.POST("/napcat_qq", services.Send_napcat_qq)
	}
	cron := router.Group("/cron")
	{
		// 设置定时任务
		cron.POST("/set", services.Cron_set)

		// 关闭定时任务
		// cron.GET("/close", cron_close)

		// 开启定时任务

		// 更新定时任务
		cron.GET("/update", services.Cron_update)

		// 删除定时任务
		cron.GET("/delete", services.Cron_delete)
		// 获取所有定时任务
		cron.GET("/list", services.Cron_list)
	}
	// 获取发送消息记录
	sendlog := router.Group("/sendlog")
	{
		sendlog.GET("/list", services.Sendlog_list)
		sendlog.GET("/delete", services.Sendlog_delete)
	}

	// 登录认证
	weblogin := router.Group("/user")
	{
		weblogin.POST("/login", services.User_login)
	}

	router.Run(fmt.Sprintf("%s:%d", apiConfig.ApiHost, apiConfig.ApiPort))
}
