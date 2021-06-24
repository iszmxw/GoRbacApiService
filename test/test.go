package main

import (
	"fmt"
	"io/ioutil"
	"net/smtp"
	"strings"
)

func SendToMail(user, sendUserName, password, host, to, subject, body, mailType string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	fmt.Println(fmt.Sprintf("%v", auth))
	var contentType string
	if mailType == "html" {
		contentType = "Content-Type: text/" + mailType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + sendUserName + "<" + user + ">" + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}

func main() {
	user := "mail@54zm.com"
	password := "zeuknxgyynudbeec"
	host := "smtp.qq.com:587"
	to := "442246396@qq.com"

	subject := "使用Golang发送邮件"

	html, e := ioutil.ReadFile("theme/001.html") //这个就是读取你的html
	if e != nil {
		panic(e)
	}
	body := string(html)
	sendUserName := "GOLANG SEND MAIL" //发送邮件的人名称
	fmt.Println("send email")
	err := SendToMail(user, sendUserName, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err.Error())
	} else {
		fmt.Println("Send mail success!")
	}

}
