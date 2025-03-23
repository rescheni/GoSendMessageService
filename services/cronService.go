package services

import (
	basic "GoMessageService/Basic"
	"GoMessageService/database"
	log "GoMessageService/log"
	"GoMessageService/sendserver"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

// CronRequest 定时任务请求结构
type CronRequest struct {
	ApiKey   string `json:"api_key" binding:"required"`
	CronExpr string `json:"cron_expr" binding:"required"`
	EntryID  string `json:"entry_id,omitempty"`
	Message  string `json:"message" binding:"required"`
	Title    string `json:"title,omitempty"`
	ToUser   string `json:"to_user,omitempty"`
	TaskType string `json:"task_type" binding:"required"` // 任务类型：wxpusher, dingding, server_jiang, email, feishu, napcat_qq
}

// Cron_list 获取所有定时任务列表
func Cron_list(c *gin.Context) {
	apiKey := c.Query("api_key")

	// 验证 API Key
	cfg := basic.LoadConfig()
	if apiKey != cfg.Api.ApiKey {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
		return
	}

	// tasks := basic.ListCronTasks()
	// c.JSON(http.StatusOK, gin.H{"tasks": tasks})
	tasks := basic.ListCronTasks()
	c.IndentedJSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})

}

// Cron_set 设置定时任务
func Cron_set(c *gin.Context) {

	// 解析请求体
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

	var inCron database.Cron
	// 检查是否存在重复的 EntryID

	inCron.CronExpr = req.CronExpr // 任务表达式
	inCron.EntryID = req.EntryID   // 任务 ID
	inCron.Message = req.Message   // 消息内容
	inCron.Title = req.Title       // 消息标题
	inCron.TaskType = req.TaskType // 任务类型
	// inCron.Status = req.IsOpen     // 任务状态
	inCron.ApiKey = req.ApiKey // API Key

	// 保存到数据库
	if !database.InsertCron(&inCron) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save cron task -> sqlite"})
		log.Logger.Error("Failed to save cron task -> sqlite")
		return
	}

	switch req.TaskType {
	case "wxpusher":
		err = basic.SetCronTask(req.CronExpr, func() {
			sendserver.SendWxPusher(req.Title, req.Message)
		})
	case "dingding":
		err = basic.SetCronTask(req.CronExpr, func() {
			sendserver.SendDing(req.Title, req.Message)
		})
	case "server_jiang":
		err = basic.SetCronTask(req.CronExpr, func() {
			sendserver.SendServerJiang(req.Title, req.Message)
		})
	case "email":
		err = basic.SetCronTask(req.CronExpr, func() {
			toUser := []string{cfg.Email.EmailAddress}
			if len(req.ToUser) > 0 {
				toUser = []string{req.ToUser}
			}
			sendserver.SendEmail(toUser, req.Title, req.Message)
		})
	case "feishu":
		err = basic.SetCronTask(req.CronExpr, func() {
			sendserver.SendFeiShu(req.Title, req.Message)
		})
	case "napcat_qq":
		err = basic.SetCronTask(req.CronExpr, func() {
			toUser := basic.LoadConfig().Napcat.NapcatQQ
			sendserver.SendQQPrivateMsg(req.Message, toUser)
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

// TODO
func Corn_close(c *gin.Context) {

}

func Cron_update(c *gin.Context) {
	// 解析请求体
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
	// 更新定时任务
	entryID, err := strconv.Atoi(req.EntryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid EntryID format"})
		return
	}
	// 根据任务类型更新不同的定时任务
	switch req.TaskType {
	case "wxpusher":
		if basic.UpdateCronTask(cron.EntryID(entryID), req.CronExpr, func() {
			sendserver.SendWxPusher(req.Title, req.Message)
		}) {
			c.JSON(http.StatusOK, gin.H{"message": "Cron task updated successfully"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cron task not found"})
		}
	case "dingding":
		if basic.UpdateCronTask(cron.EntryID(entryID), req.CronExpr, func() {
			sendserver.SendDing(req.Title, req.Message)
		}) {
			c.JSON(http.StatusOK, gin.H{"message": "Cron task updated successfully"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cron task not found"})
		}
	case "server_jiang":
		if basic.UpdateCronTask(cron.EntryID(entryID), req.CronExpr, func() {
			sendserver.SendServerJiang(req.Title, req.Message)
		}) {
			c.JSON(http.StatusOK, gin.H{"message": "Cron task updated successfully"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cron task not found"})
		}
	case "email":
		if basic.UpdateCronTask(cron.EntryID(entryID), req.CronExpr, func() {
			sendserver.SendEmail([]string{cfg.Email.EmailAddress}, req.Title, req.Message)
		}) {
			c.JSON(http.StatusOK, gin.H{"message": "Cron task updated successfully"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cron task not found"})
		}
	case "feishu":
		if basic.UpdateCronTask(cron.EntryID(entryID), req.CronExpr, func() {
			sendserver.SendFeiShu(req.Title, req.Message)
		}) {
			c.JSON(http.StatusOK, gin.H{"message": "Cron task updated successfully"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cron task not found"})
		}
	case "napcat_qq":
		if basic.UpdateCronTask(cron.EntryID(entryID), req.CronExpr, func() {
			sendserver.SendQQPrivateMsg(req.Message, "1413024330")
		}) {
			c.JSON(http.StatusOK, gin.H{"message": "Cron task updated successfully"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cron task not found"})
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task type"})
		return
	}

}

func Cron_delete(c *gin.Context) {
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

	database.DeleteCron(entryIDStr)

	if basic.DeleteCronTask(cron.EntryID(entryID)) {
		c.JSON(http.StatusOK, gin.H{"message": "Cron task deleted successfully"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cron task not found"})
	}
}
