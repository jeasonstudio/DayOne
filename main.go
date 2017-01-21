package main

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

const (
	SERVICE  = "smtp.exmail.qq.com"
	PORT     = 25
	ALIAS    = "每天一题"
	SUBJECT  = "每天一题"
	USER     = "mailbyjeason@jeasonstudio.cn" //发送邮件的邮箱
	PASSWORD = "Admin12345"                   //发送邮件邮箱的密码
)

func sendEmail() {

	m := gomail.NewMessage()
	m.SetAddressHeader("From", USER, "Jeason") // 发件人
	m.SetHeader("To", "zjtong3576@sina.com", "me@jeasonstudio.cn")

	m.SetHeader("Subject", ALIAS)
	m.SetBody("text/html", "Hello <a href = \"http://blog.csdn.net/liang19890820\">一去丶二三里</a>") // 正文

	fmt.Println("Ready to sendEmail...")

	d := gomail.NewPlainDialer(SERVICE, PORT, USER, PASSWORD) // 发送邮件服务器、端口、发件人账号、发件人密码

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Sending err,", err)
		panic(err)
	}
	fmt.Println("Send Success!")

}

func main() {
	sendEmail()
	// fmt.Println(a)
}
