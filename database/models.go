package database

// type User struct {
// 	ID       uint   `json:"id" gorm:"primaryKey"`
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// 	Email    string `json:"email"`
// }

type Cron struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	ApiKey   string `json:"api" binding:"required"`
	CronExpr string `json:"cron_expr" binding:"required"`
	EntryID  string `json:"entry_id,omitempty"`
	Message  string `json:"message" binding:"required"`
	Title    string `json:"title,omitempty"`
	TaskType string `json:"task_type" binding:"required"` // 任务类型：wxpusher, dingding, server_jiang, email, feishu, napcat_qq
	// 任务状态：0-未启动，1-已启动
	Status int `json:"status"`
	// 所属用户
	ByUserId int `json:"by_user_id"`
}
