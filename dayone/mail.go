package dayone

import (
	f "fmt"

	"gopkg.in/gomail.v2"
)

// const HOST 数据库相关配置
const (
	HOST     = "smtp.exmail.qq.com"
	PORT     = 25
	SUBJECT  = "每日一题"
	USER     = "mailbyjeason@jeasonstudio.cn"
	PASSWORD = "Admin12345"
)

// SendEmail 发送邮件
func SendEmail(innerHTML string) {

	m := gomail.NewMessage()
	m.SetAddressHeader("From", USER, "Jeason")
	m.SetHeader("To",
		// m.FormatAddress("zjtong3576@sina.com", "Jeason"),
		// m.FormatAddress("me@jeasonstudio.cn", "Jeason"),
		m.FormatAddress("748807384@qq.com", " 傻了吧唧的贾同学"),
	)

	m.SetHeader("Subject", SUBJECT)
	m.SetBody("text/html", innerHTML)

	f.Println("Send Ready...")

	d := gomail.NewPlainDialer(HOST, PORT, USER, PASSWORD)
	// 发送邮件服务器、端口、发件人账号、发件人密码

	if err := d.DialAndSend(m); err != nil {
		f.Println("Send Fail with,", err)
		panic(err)
	}
	f.Println("Send Email Success!!")
}
