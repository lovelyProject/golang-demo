package mail

import (
	"crypto/tls"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendMail(login, password, host, hash string) {
	e := email.NewEmail()
	e.From = "My App <" + login + ">" // именно так
	e.To = []string{"edward.hardwork2000@gmail.com"}
	e.Bcc = []string{"edward.hardwork2000@gmail.com"}
	e.Cc = []string{"edward.hardwork2000@gmail.com"}
	e.Subject = "Awesome Subject"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte(hash)

	auth := smtp.PlainAuth("", login, password, host)

	err := e.SendWithTLS(
		host+":465",
		auth,
		&tls.Config{ServerName: host},
	)
	if err != nil {
		log.Printf("send mail failed: %v", err)
	} else {
		log.Println("mail sent successfully")
	}
}
