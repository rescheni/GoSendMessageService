package sendserver

import (
	basic "GoMessageService/Basic"
	"context"
	"encoding/json"
	"fmt"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

type TextMessage struct {
	Text string `json:"text"`
}

func FeiShuSend(text string, desp string) {
	// 创建 Client
	// client := lark.NewClient("YOUR_APP_ID", "YOUR_APP_SECRET")

	cfg := basic.LoadConfig()

	appId := cfg.Feishu.FeishuAppId
	appSecret := cfg.Feishu.FeishuAppSecret
	FeishuUserId := cfg.Feishu.FeishuUserId

	client := lark.NewClient(appId, appSecret)

	// 使用 json 库处理消息内容
	msg := TextMessage{Text: "<b>" + desp + "</b>"}
	content, err := json.Marshal(msg)
	if err != nil {
		fmt.Printf("JSON 编码错误: %v\n", err)
		return
	}

	fmt.Printf("发送的内容: %s\n", content)
	fmt.Println("--------------------------------")

	// 创建请求对象
	req := larkim.NewCreateMessageReqBuilder().
		ReceiveIdType("open_id").
		Body(larkim.NewCreateMessageReqBodyBuilder().
			ReceiveId(FeishuUserId).
			MsgType("text").
			Content(string(content)).
			Build()).
		Build()

	// 发起请求
	resp, err := client.Im.Message.Create(context.Background(), req)

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Printf("logId: %s, error response: \n%s", resp.RequestId, larkcore.Prettify(resp.CodeError))
		return
	}

	// 业务处理
	fmt.Println(larkcore.Prettify(resp))
}
