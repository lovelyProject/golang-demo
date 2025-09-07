package main

import (
	"go/adv-example/configs"
	"go/adv-example/internal/handler"
	"net/http"
)

func main() {
	conf := configs.NewConfig()
	router := http.NewServeMux()
	handler.NewHandler(router)
}
