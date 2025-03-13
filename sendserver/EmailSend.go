package sendserver

import (
	basic "GoMessageService/Basic"
	"fmt"

	"crypto/tls"

	"gopkg.in/gomail.v2"
)

// EmailSend 用于发送邮件
func EmailSend(sendTo []string) {
	emailTest(sendTo)
}

// emailTest 发送测试邮件
func emailTest(sendTo []string) {
	// 加载配置
	cfg := basic.LoadConfig()

	host := cfg.Email.SmtpServer
	port := cfg.Email.SmtpPort
	userName := cfg.Email.Username
	password := cfg.Email.AuthCode

	// 创建 gomail.NewMessage() 邮件对象
	m := gomail.NewMessage()
	m.SetHeader("From", userName)
	m.SetHeader("To", sendTo...)
	m.SetHeader("Subject", "Email Send Test")
	m.SetBody("text/plain", "Text Body is, of course, supported!")
	m.SetBody("text/html", "<h1>This a test email</h1>")

	// 创建一个新的 gomail.Dialer 对象，它支持 SSL/TLS
	dialer := gomail.NewDialer(host, port, userName, password)
	dialer.SSL = true
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// 打印连接信息
	fmt.Printf("Attempting to connect to SMTP server: %s:%d\n", host, port)
	fmt.Printf("Username: %s\n", userName)

	// 使用 Dialer 发送邮件
	if err := dialer.DialAndSend(m); err != nil {
		fmt.Printf("Error sending email: %v\n", err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}
