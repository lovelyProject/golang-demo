package handler

import (
	"fmt"
	"net/http"
)

type SmtpHandler struct {
}

func NewHandler(router *http.ServeMux) {
	handler := &SmtpHandler{}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("/verify/{hash}", handler.Send())
}

func (handler *SmtpHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Send")
	}
}

func (handler *SmtpHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("verify")
	}
}
