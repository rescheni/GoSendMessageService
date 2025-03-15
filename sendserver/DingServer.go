package sendserver

import (
	basic "GoMessageService/Basic"
	"bytes"
	"encoding/json"
	"net/http"
)

type Message struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

func DingServer(content string) {
	cfg := basic.LoadConfig()
	webhookURL := "https://oapi.dingtalk.com/robot/send?access_token=" + cfg.DingDing.AccessToken

	message := Message{
		MsgType: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: content,
		},
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 处理响应
}
