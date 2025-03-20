package services

type DataTime struct {
	ApiKey  string   `json:"api_key" binding:"required"`
	Message string   `json:"message" binding:"required"`
	ToUser  []string `json:"to_user,omitempty"`
	Title   string   `json:"title,omitempty"`
}
