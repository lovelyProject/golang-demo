package main

import (
	"math/rand"
	"net/http"
	"strconv"
)

type RandHander struct{}

func NewRandHandler(router *http.ServeMux) {
	handler := &RandHander{}
	router.HandleFunc("/rand", handler.RandNumber())
}

func (r *RandHander) RandNumber() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(strconv.Itoa(rand.Intn(7))))
	}
}
