package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func sendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")

	auth := smtp.PlainAuth("", user, password, hp[0])

	var contentType string
	if mailtype == "html" {
		contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}

func main() {
	user := "hank@hankbook.cn" // 邮箱账号
	password := "****"         // 邮箱授权密码，不是邮箱密码
	host := "smtp.qq.com:25"   // smtp地址
	to := "123456@qq.com"      // 发送邮箱
	subject := "世界是美好的"

	body := "生日快乐"
	fmt.Println("send email")
	err := sendToMail(user, password, host, to, subject, body, "")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}

}
