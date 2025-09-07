package main

import (
	"fmt"
	"go/adv-example/configs"
	"go/adv-example/internal/handler"
	"net/http"
)

func main() {
	conf := configs.NewConfig()
	router := http.NewServeMux()
	handler.NewHandler(router)

	server := http.Server{
		Addr:    conf.Port,
		Handler: router,
	}

	fmt.Println(conf.Port)

	server.ListenAndServe()
}
