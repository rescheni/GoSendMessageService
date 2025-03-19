package services

import (
	basic "GoMessageService/Basic"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User_login 用户登录
func User_login(c *gin.Context) {
	apiToken := c.Query("api_key") // 修改为 api_token

	cfg := basic.LoadConfig()

	if apiToken != cfg.Api.ApiKey { // 验证 api_token
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API_KEY"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"}) // 登录成功响应
}
