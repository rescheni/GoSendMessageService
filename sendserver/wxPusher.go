package sendserver

import (
	basic "GoMessageService/Basic"
	"fmt"

	"github.com/wxpusher/wxpusher-sdk-go"
	"github.com/wxpusher/wxpusher-sdk-go/model"
)

// SendWxPusher 发送消息给指定用户
func SendWxPusher(title, content string) {
	cf := basic.LoadConfig()
	appToken := cf.WxPusher.WxPusherKey
	uId := cf.WxPusher.DefaultUid

	// 创建消息并设置内容
	msg := model.NewMessage(appToken).
		SetContent(fmt.Sprintf("【%s】\n%s", title, content)).
		SetSummary(title).
		AddUId(uId)

	// 发送消息
	msgArr, err := wxpusher.SendMessage(msg)
	if err != nil {
		fmt.Printf("发送消息失败: %v\n", err)
		return
	}

	if len(msgArr) > 0 {
		mid := msgArr[0].MessageId
		WxPusherState(mid)
		fmt.Printf("消息发送成功，消息ID: %d\n", mid)
	}
}

// WxPusherState 查询消息状态
func WxPusherState(messageId int) {
	status, err := wxpusher.QueryMessageStatus(messageId)
	if err != nil {
		fmt.Printf("查询消息状态失败: %v\n", err)
		return
	}
	fmt.Printf("消息状态: %+v\n", status)
}

// WxPusherCreateQRCode 创建二维码
func WxPusherCreateQRCode() {
	cf := basic.LoadConfig()
	appToken := cf.WxPusher.WxPusherKey

	qrcode := model.Qrcode{
		AppToken:  appToken,
		Extra:     "WxPusher",
		ValidTime: 3600, // 二维码有效期1小时
	}

	qrcodeResp, err := wxpusher.CreateQrcode(&qrcode)
	if err != nil {
		fmt.Printf("创建二维码失败: %v\n", err)
		return
	}

	fmt.Printf("二维码创建成功:\n")
	fmt.Printf("- 二维码地址: %s\n", qrcodeResp.Url)
	fmt.Printf("- 二维码图片: %s\n", qrcodeResp.ShortUrl)
}

// WxPusherUserList 获取用户列表
func WxPusherUserList() {
	cf := basic.LoadConfig()
	appToken := cf.WxPusher.WxPusherKey

	result, err := wxpusher.QueryWxUser(appToken, 1, 20)
	if err != nil {
		fmt.Printf("获取用户列表失败: %v\n", err)
		return
	}

	fmt.Printf("用户列表:\n")
	for i, user := range result.Records {
		fmt.Printf("%d. 用户ID: %s, 昵称: %s\n", i+1, user.UId, user.NickName)
	}
}

func WxPusherAddUserList() {

}
