package handler

import (
	"fmt"
	"go/adv-example/pkg/hash"
	request "go/adv-example/pkg/req"
	"go/adv-example/pkg/res"
	sendMail "go/adv-example/pkg/smtp"
	"net/http"
	_ "net/smtp"
	_ "os"

	"go/adv-example/configs"

	_ "github.com/jordan-wright/email"
)

type SmtpHandler struct {
	Config configs.SmtpConfig
}

func NewHandler(router *http.ServeMux, deps configs.SmtpConfig) {
	handler := &SmtpHandler{
		Config: deps,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("/verify/{hash}", handler.Verify())
	router.HandleFunc("POST /register", handler.Register())
}

func (handler *SmtpHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, _ := request.HandleBody[SendRequest](w, req)
		fmt.Println(body, handler.Config)
	}
}

func (handler *SmtpHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, _ := request.HandleBody[SendRequest](w, req)
		fmt.Println(body)
	}
}

func (handler *SmtpHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, _ := request.HandleBody[RegisterRequest](w, req)
		newHash, _ := hash.GenerateHash(32)
		sendMail.SendMail(handler.Config.Email, handler.Config.Password, handler.Config.Server, newHash)
		res.Json(w, newHash, 200)
		fmt.Println(body)
	}
}
