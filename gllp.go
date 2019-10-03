package main

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	user := "xxxx@xx.com"
	password := "xxxxxx"
	host := "smtp.qq.com:25"
	to := "xxxx@qq.com"

	subject := "使用Golang发送邮件"

	body := `<html><body><h3>"Test golang send to email"</h3></body></html>`
	fmt.Println("send email")
	err := SendMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}

type mailConf struct {
	User     string
	Password string
	Host     string
	To       string
	Title    string
	Body     string
	Type     string
}

func SendMail(user, password, host, to, subject, body, mailtype string) error {
	fmt.Println("Sending...")
	conf := mailConf{User: user, Password: password, Host: host, To: to, Title: subject, Body: body, Type: mailtype}
	return doSend(conf)
}

func doSend(conf mailConf) error {
	hp := strings.Split(conf.Host, ":")
	auth := smtp.PlainAuth("", conf.User, conf.Password, hp[0])
	var contentType string
	if conf.Type == "html" {
		contentType = "Content-Type: text/" + conf.Type + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain; charset=UTF-8"
	}
	msg := []byte("To: " + conf.To + "\r\nFrom: " + conf.User + ">\r\nSubject: " + "\r\n" + contentType + "\r\n\r\n" + conf.Body)
	send_to := strings.Split(conf.To, ";")
	err := smtp.SendMail(conf.Host, auth, conf.User, send_to, msg)
	return err
}
