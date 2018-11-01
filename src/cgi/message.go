package cgi

import (
	"fmt"
	"net/smtp"
	"strings"
)

type mailConf struct {
	User string
	Password string
	Host string
	To string
	Title string
	Body string
	Type string
}


func SendMail(user, password, host, to, subject, body, mailtype string) error {
	fmt.Println("Sending...")
	conf := mailConf{User:user,Password:password,Host:host,To:to,Title:subject,Body:body,Type:mailtype}
	return doSend(conf)
}

func doSend(conf mailConf) error {
	hp :=strings.Split(conf.Host,":")
	auth := smtp.PlainAuth("",conf.User,conf.Password,hp[0])
	var contentType string
	if conf.Type=="html" {
		contentType = "Content-Type: text/"+conf.Type+"; charset=UTF-8"
	}else {
		contentType = "Content-Type: text/plain; charset=UTF-8"
	}
	msg := []byte("To: "+conf.To+"\r\nFrom: "+conf.User+">\r\nSubject: " + "\r\n" + contentType + "\r\n\r\n" + conf.Body)
	send_to := strings.Split(conf.To, ";")
	err := smtp.SendMail(conf.Host, auth, conf.User, send_to, msg)
	return err
}
