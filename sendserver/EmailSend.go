package sendserver

import (
	basic "GoMessageService/Basic"
	log "GoMessageService/log"
	"fmt"

	"crypto/tls"

	"gopkg.in/gomail.v2"
)

// SendEmail 用于发送邮件
func SendEmail(sendTo []string, title string, message string) {
	emailTest(sendTo, title, message)
}

// emailTest 发送测试邮件
func emailTest(sendTo []string, message string, subject string) {
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
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", message)
	m.SetBody("text/html", message)

	// 创建一个新的 gomail.Dialer 对象，它支持 SSL/TLS
	dialer := gomail.NewDialer(host, port, userName, password)
	dialer.SSL = true
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// 打印连接信息
	// fmt.Printf("Attempting to connect to SMTP server: %s:%d\n", host, port)
	log.Logger.Info(fmt.Sprintf("Attempting to connect to SMTP server: %s:%d\n", host, port))
	// fmt.Printf("Username: %s\n", userName)
	log.Logger.Info(fmt.Sprintf("Username: %s\n", userName))

	// 使用 Dialer 发送邮件
	if err := dialer.DialAndSend(m); err != nil {
		// fmt.Printf("Error sending email: %v\n", err)
		log.Logger.Error(fmt.Sprintf("Error sending email: %v\n", err))
	} else {
		// fmt.Println("Email sent successfully!")
		log.Logger.Info("Email sent successfully!")
	}
}
