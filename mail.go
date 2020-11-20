package Mailsender

import(
	"gopkg.in/gomail.v2"
)

func Sendmail(a string){
	m := gomail.NewMessage()
	// 发邮件的地址
	m.SetHeader("From", "shenboyu2020@163.com")
	// 给谁发送，支持多个账号
	m.SetHeader("To", "1720648723@qq.com")
	// 抄送谁
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	// 邮件标题
	m.SetHeader("Subject", "芜湖起飞from沈博宇")
	// 邮件正文，支持 html
	m.SetBody("text/html", a)
	// 附件
	//m.Attach("/home/Alex/lolcat.jpg")
	// stmp服务，端口号，发送邮件账号，发送账号密码
	d := gomail.NewDialer("smtp.163.com", 25, "shenboyu2020@163.com", "WOEVMSPVCYASVDNS")
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}