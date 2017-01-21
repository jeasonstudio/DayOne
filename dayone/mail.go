package dayone

import (
	"gopkg.in/gomail.v2"
)

const (
	HOST       = "smtp.exmail.qq.com"
	SERVERADDR = "smtp.exmail.qq.com:25"
	ALIAS      = "每天一题"
	SUBJECT    = "每天一题"
	USER       = "mailbyjeason@jeasonstudio.cn" //发送邮件的邮箱
	PASSWORD   = "Admin12345"                   //发送邮件邮箱的密码
)

func sendEmail() string {

	m := gomail.NewMessage()
	m.SetAddressHeader("From", USER, "一去、二三里") // 发件人
	m.SetHeader("To",                          // 收件人
		m.FormatAddress("zjtong3576@sina.com", "乔峰"),
	)

	m.SetHeader("这是", "啥啊")                                                                     // 主题
	m.SetBody("text/html", "Hello <a href = \"http://blog.csdn.net/liang19890820\">一去丶二三里</a>") // 正文

	d := gomail.NewPlainDialer(HOST, 25, USER, PASSWORD) // 发送邮件服务器、端口、发件人账号、发件人密码

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return "ok"
}
