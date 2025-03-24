package sendserver

import (
	basic "GoMessageService/Basic"
	log "GoMessageService/log"
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

func SendDing(title, content string) {
	cfg := basic.LoadConfig()
	webhookURL := "https://oapi.dingtalk.com/robot/send?access_token=" + cfg.DingDing.AccessToken

	message := Message{
		MsgType: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: title + "\n" + content,
		},
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Logger.Error("json.Marshal failed ")
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Logger.Error("http.Post failed ")
	}
	defer resp.Body.Close()

}
