package mail

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func SendMail(login, password, host, hash, emailTo string) error {
	e := email.NewEmail()
	e.From = "My App <" + login + ">" // именно так
	e.To = []string{"edward.hardwork2000@gmail.com"}
	e.Bcc = []string{"edward.hardwork2000@gmail.com"}
	e.Cc = []string{"edward.hardwork2000@gmail.com"}
	e.Subject = "Awesome Subject"
	e.Text = []byte("Привет! Перейди по ссылке для подтверждения: http://localhost:8081/verify/" + hash)

	e.HTML = []byte(fmt.Sprintf(
		`<p>Привет!</p>
    <p>Перейди по ссылке для подтверждения:</p>
    <a href="http://localhost:8081/verify/%s">Подтвердить email</a>`, hash))

	auth := smtp.PlainAuth("", login, password, host)

	err := e.SendWithTLS(
		host+":465",
		auth,
		&tls.Config{ServerName: host},
	)
	if err != nil {
		return err
	} else {
		return nil
	}
}
