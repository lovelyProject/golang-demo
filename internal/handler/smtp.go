package handler

import (
	"fmt"
	_ "github.com/jordan-wright/email"
	"go/adv-example/configs"
	"go/adv-example/pkg/file"
	"go/adv-example/pkg/hash"
	request "go/adv-example/pkg/req"
	"go/adv-example/pkg/res"
	sendMail "go/adv-example/pkg/smtp"
	"net/http"
	_ "net/smtp"
	_ "os"
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
		newHash, _ := hash.GenerateHash(32)
		err := sendMail.SendMail(handler.Config.Email, handler.Config.Password, handler.Config.Server, newHash, body.Email)
		if err != nil {
			fmt.Println(err.Error())
		}
		file.SaveInFile("hash.json", body.Email, newHash)
		res.Json(w, newHash, 200)

	}
}

func (handler *SmtpHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		pathHash := req.PathValue("hash")
		isValid := file.CompareText("hash.json", pathHash)
		if !isValid {
			file.DeleteFile("hash.json")
			res.Json(w, "invalid hash", 400)
			return
		}

		fmt.Println("verified success")
		file.DeleteFile("hash.json")
		res.Json(w, "verified success", 200)
		// body, _ := request.HandleBody[SendRequest](w, req)
		// fmt.Println(body)
	}
}

func (handler *SmtpHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, _ := request.HandleBody[RegisterRequest](w, req)
		fmt.Println(body)
	}
}
