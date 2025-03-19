package services

// 发送消息服务

import (
	basic "GoMessageService/Basic"
	"GoMessageService/sendserver"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MessageRequest 消息请求结构
type MessageRequest struct {
	ApiKey  string `json:"api_key" binding:"required"`
	Message string `json:"message" binding:"required"`
	ToUser  string `json:"to_user,omitempty"`
	Title   string `json:"title,omitempty"`
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

	// 发送消息
	sendserver.SendEmail([]string{cfg.Email.EmailAddress}, req.Title, req.Message)
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

	// 发送消息
	sendserver.SendQQPrivateMsg(req.Message, "1413024330")
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}
