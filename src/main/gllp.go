package main

import (
	"cgi"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	str := cgi.Get()
	fmt.Println(str)
	user := "xxxx@xx.com"
	password := "xxxxxx"
	host := "smtp.qq.com:25"
	to := "xxxx@qq.com"

	subject := "使用Golang发送邮件"

	body := `<html><body><h3>"Test golang send to email"</h3></body></html>`
	fmt.Println("send email")
	err := cgi.SendMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}
