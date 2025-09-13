package service

import (
	"go/adv-example/pkg/res"
	"net/http"
)

type Handler struct{}

func NewHandler(router *http.ServeMux) {
	handler := &Handler{}
	router.HandleFunc("/", handler.Test())
}

func (h *Handler) Test() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res.Json(w, http.StatusOK, "Hello, World!")
	}
}
