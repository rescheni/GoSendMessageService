package sendserver

import (
	basic "GoMessageService/Basic"
	log "GoMessageService/log"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func SendQQPrivateMsg(message string, userID string) {

	cfg := basic.LoadConfig()

	url := cfg.Napcat.NapcatUrl + "/send_private_msg"
	method := "POST"

	print(message)

	type MessageData struct {
		Text string `json:"text"`
	}

	type Message struct {
		Type string      `json:"type"`
		Data MessageData `json:"data"`
	}

	type PayloadStruct struct {
		UserID  string    `json:"user_id"`
		Message []Message `json:"message"`
	}

	payloadData := PayloadStruct{
		UserID: userID,
		Message: []Message{
			{
				Type: "text",
				Data: MessageData{
					Text: message,
				},
			},
		},
	}

	jsonData, err := json.Marshal(payloadData)
	if err != nil {
		// fmt.Println("JSON编码错误:", err)
		log.Logger.Error("JSON编码错误:", err)
		return
	}

	payload := strings.NewReader(string(jsonData))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		// fmt.Println(err)
		log.Logger.Error(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+cfg.Napcat.NapcatToken)

	res, err := client.Do(req)
	if err != nil {
		// fmt.Println(err)
		log.Logger.Error(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// fmt.Println(err)
		log.Logger.Error(err)
		return
	}
	// fmt.Println(string(body))
	log.Logger.Info(string(body))
}

func SendQQGroupMessage(message string, groupID string) {

	cfg := basic.LoadConfig()

	url := cfg.Napcat.NapcatUrl + "/send_group_msg"
	method := "POST"

	type MessageData struct {
		Text string `json:"text"`
	}

	type Message struct {
		Type string      `json:"type"`
		Data MessageData `json:"data"`
	}

	type PayloadStruct struct {
		GroupID string    `json:"group_id"`
		Message []Message `json:"message"`
	}

	payloadData := PayloadStruct{
		GroupID: groupID,
		Message: []Message{
			{
				Type: "text",
				Data: MessageData{
					Text: message,
				},
			},
		},
	}

	jsonData, err := json.Marshal(payloadData)
	if err != nil {
		// fmt.Println("JSON编码错误:", err)
		log.Logger.Error("JSON编码错误:", err)
		return
	}

	payload := strings.NewReader(string(jsonData))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		// fmt.Println(err)
		log.Logger.Error(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+cfg.Napcat.NapcatToken)

	res, err := client.Do(req)
	if err != nil {
		// fmt.Println(err)
		log.Logger.Error(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// fmt.Println(err)
		log.Logger.Error(err)
		return
	}
	// fmt.Println(string(body))
	log.Logger.Info(string(body))
}
