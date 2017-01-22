package dayone

import (
	f "fmt"

	"gopkg.in/gomail.v2"
)

const (
	HOST     = "smtp.exmail.qq.com" //HOST 邮件服务器
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
		m.FormatAddress("zjtong3576@sina.com", "Jeason"),
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
