package api

import (
	basic "GoMessageService/Basic"
	"GoMessageService/sendserver"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

// MessageRequest 消息请求结构
type MessageRequest struct {
	ApiKey  string `json:"api_key" binding:"required"`
	Message string `json:"message" binding:"required"`
	Title   string `json:"title,omitempty"`
}

// CronRequest 定时任务请求结构
type CronRequest struct {
	ApiKey   string `json:"api_key" binding:"required"`
	CronExpr string `json:"cron_expr" binding:"required"`
	EntryID  string `json:"entry_id,omitempty"`
	Message  string `json:"message" binding:"required"`
	Title    string `json:"title,omitempty"`
	IsOpen   bool   `json:"is_open"`
	TaskType string `json:"task_type" binding:"required"` // 任务类型：wxpusher, dingding, server_jiang, email, feishu, napcat_qq
}

func APIStart() {
	cfg := basic.LoadConfig()
	apiConfig := cfg.Api

	router := gin.Default()

	send := router.Group("/send")
	{
		// 发送消息
		send.POST("/wxpusher", send_wxpusher)
		send.POST("/dingding", send_dingding)
		send.POST("/server_jiang", send_server_jiang)
		send.POST("/email", send_email)
		send.POST("/feishu", send_feishu)
		send.POST("/napcat_qq", send_napcat_qq)
	}

	cron := router.Group("/cron")
	{
		// 设置定时任务
		cron.POST("/set", cron_set)

		// 关闭定时任务
		// cron.GET("/close", cron_close)

		// 删除定时任务
		cron.GET("/delete", cron_delete)
		// 获取所有定时任务
		cron.GET("/list", cron_list)
	}

	// myuser := router.Group("/user")
	// {
	// 	myuser.POST("/login", user_login)
	// 	myuser.GET("/status", user_status)
	// 	myuser.POST("/register", user_register)
	// 	myuser.POST("/logout", user_logout)
	// }

	router.Run(fmt.Sprintf("%s:%d", apiConfig.ApiHost, apiConfig.ApiPort))
}

func send_wxpusher(c *gin.Context) {
	var req MessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证 API Key
	cfg := basic.LoadConfig()
	if req.ApiKey != cfg.Api.ApiKey {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
		return
	}

	// 发送消息
	sendserver.SendWxPusher(req.Title, req.Message)
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func send_dingding(c *gin.Context) {
	var req MessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证 API Key
	cfg := basic.LoadConfig()
	if req.ApiKey != cfg.Api.ApiKey {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
		return
	}

	// 发送消息
	sendserver.DingSend(req.Message)
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func send_server_jiang(c *gin.Context) {
	var req MessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证 API Key
	cfg := basic.LoadConfig()
	if req.ApiKey != cfg.Api.ApiKey {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
		return
	}

	// 发送消息
	sendserver.ServerJiang(req.Title, req.Message)
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func send_email(c *gin.Context) {
	var req MessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证 API Key
	cfg := basic.LoadConfig()
	if req.ApiKey != cfg.Api.ApiKey {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
		return
	}

	// 发送消息
	sendserver.EmailSend([]string{cfg.Email.EmailAddress}, req.Message, req.Title)
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func send_feishu(c *gin.Context) {
	var req MessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证 API Key
	cfg := basic.LoadConfig()
	if req.ApiKey != cfg.Api.ApiKey {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
		return
	}

	// 发送消息
	sendserver.FeiShuSend(req.Title, req.Message)
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func send_napcat_qq(c *gin.Context) {
	var req MessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证 API Key
	cfg := basic.LoadConfig()
	if req.ApiKey != cfg.Api.ApiKey {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
		return
	}

	// 发送消息
	sendserver.Send_private_msg(req.Message, "1413024330")
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func cron_set(c *gin.Context) {
	var req CronRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证 API Key
	cfg := basic.LoadConfig()
	if req.ApiKey != cfg.Api.ApiKey {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
		return
	}

	// 根据任务类型设置不同的定时任务
	var err error

	if req.IsOpen == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid IsOpen format"})
		return
	}

	switch req.TaskType {
	case "wxpusher":
		err = basic.SetCronTask(req.CronExpr, func() {
			sendserver.SendWxPusher(req.Title, req.Message)
		})
	case "dingding":
		err = basic.SetCronTask(req.CronExpr, func() {
			sendserver.DingSend(req.Message)
		})
	case "server_jiang":
		err = basic.SetCronTask(req.CronExpr, func() {
			sendserver.ServerJiang(req.Title, req.Message)
		})
	case "email":
		err = basic.SetCronTask(req.CronExpr, func() {
			sendserver.EmailSend([]string{cfg.Email.EmailAddress}, req.Message, req.Title)
		})
	case "feishu":
		err = basic.SetCronTask(req.CronExpr, func() {
			sendserver.FeiShuSend(req.Title, req.Message)
		})
	case "napcat_qq":
		err = basic.SetCronTask(req.CronExpr, func() {
			sendserver.Send_private_msg(req.Message, "1413024330")
		})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task type"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Cron task set successfully",
		"cron_expr": req.CronExpr,
	})
}

func corn_close(c *gin.Context) {

}

func cron_delete(c *gin.Context) {
	apiKey := c.Query("api_key")
	entryIDStr := c.Query("entryid") // 修改为小写 "entryid"

	// 验证 API Key
	cfg := basic.LoadConfig()
	if apiKey != cfg.Api.ApiKey {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
		return
	}

	// 删除定时任务
	entryID, err := strconv.Atoi(entryIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid EntryID format"})
		return
	}
	if basic.DeleteCronTask(cron.EntryID(entryID)) {
		c.JSON(http.StatusOK, gin.H{"message": "Cron task deleted successfully"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cron task not found"})
	}
}

func cron_list(c *gin.Context) {
	apiKey := c.Query("api_key")

	// 验证 API Key
	cfg := basic.LoadConfig()
	if apiKey != cfg.Api.ApiKey {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
		return
	}

	tasks := basic.ListCronTasks()
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func user_login(c *gin.Context) {

}
