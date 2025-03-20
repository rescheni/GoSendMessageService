package services

// 发送消息服务

import (
	basic "GoMessageService/Basic"
	"GoMessageService/database"
	log "GoMessageService/log"
	"GoMessageService/sendserver"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// MessageRequest 消息请求结构
type MessageRequest struct {
	ApiKey  string   `json:"api_key" binding:"required"`
	Message string   `json:"message" binding:"required"`
	ToUser  []string `json:"to_user,omitempty"`
	Title   string   `json:"title,omitempty"`
}

func Send_wxpusher(c *gin.Context) {
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
	sendtime := time.Now().Format("2006-01-02 15:04:05")
	log.Logger.Info("WxPusher 消息发送成功")
	database.SaveSendedMessage(sendtime, req.Title, req.Message, "Wxpusher", "success")

	// 发送消息
	sendserver.SendWxPusher(req.Title, req.Message)
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})

}

func Send_dingding(c *gin.Context) {
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
	sendtime := time.Now().Format("2006-01-02 15:04:05")
	log.Logger.Info("钉钉消息发送成功")
	database.SaveSendedMessage(sendtime, req.Title, req.Message, "dingding", "success")

	// 发送消息
	sendserver.SendDing(req.Title, req.Message)
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func Send_server_jiang(c *gin.Context) {
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
	sendtime := time.Now().Format("2006-01-02 15:04:05")
	log.Logger.Info("Server酱 消息发送成功")
	database.SaveSendedMessage(sendtime, req.Title, req.Message, "Server酱", "success")

	// 发送消息
	sendserver.SendServerJiang(req.Title, req.Message)
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func Send_email(c *gin.Context) {
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
	sendtime := time.Now().Format("2006-01-02 15:04:05")
	log.Logger.Info("Email消息发送成功")
	database.SaveSendedMessage(sendtime, req.Title, req.Message, "Email", "success")

	tosendUser := req.ToUser
	// 如果 ToUser 为空，则发送到配置文件中的邮箱
	if len(tosendUser) == 0 {
		tosendUser = []string{cfg.Email.EmailAddress}
	}

	// 发送消息
	sendserver.SendEmail(tosendUser, req.Title, req.Message)
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func Send_feishu(c *gin.Context) {
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
	sendtime := time.Now().Format("2006-01-02 15:04:05")
	log.Logger.Info("飞书 消息发送成功")
	database.SaveSendedMessage(sendtime, req.Title, req.Message, "飞书", "success")

	// 发送消息
	sendserver.SendFeiShu(req.Title, req.Message)
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func Send_napcat_qq(c *gin.Context) {
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
	sendtime := time.Now().Format("2006-01-02 15:04:05")
	log.Logger.Info("napcat 消息发送成功")
	database.SaveSendedMessage(sendtime, req.Title, req.Message, "napcat", "success")

	// 发送消息
	sendserver.SendQQPrivateMsg(req.Message, "1413024330")
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func Sendlog_list(c *gin.Context) {
	// 验证api key
	cfg := basic.LoadConfig()
	if cfg.Api.ApiKey != c.Query("api_key") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
		return
	}

	sendlog := database.GetSendedMessage()
	c.JSON(http.StatusOK, gin.H{"message": sendlog})
}

func Sendlog_delete(c *gin.Context) {
	// 验证api key
	cfg := basic.LoadConfig()
	if cfg.Api.ApiKey != c.Query("api_key") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
		return
	}

	sendid := c.Query("sendid")
	id, err := strconv.Atoi(sendid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sendid"})
		return
	}

	// 检查是否存在
	sendlog := database.GetSendedMessage()
	if len(sendlog) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Send log not found"})
		return
	}

	database.DeleteSendedMessage(id)
	c.JSON(http.StatusOK, gin.H{"message": "Send log deleted successfully"})
}
