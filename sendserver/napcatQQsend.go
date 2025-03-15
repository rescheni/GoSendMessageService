package sendserver

import (
	basic "GoMessageService/Basic"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Send_private_msg(message string, userID string) {

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
		fmt.Println("JSON编码错误:", err)
		return
	}

	payload := strings.NewReader(string(jsonData))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+cfg.Napcat.NapcatToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func Send_group_msg(message string, groupID string) {

	url := "/send_group_msg"
	method := "POST"

	payload := strings.NewReader(`{
    "group_id": "123456",
    "message": [
        {
            "type": "text",
            "data": {
                "text": "napcat"
            }
        }
    ]
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
