package main

import (
	"fmt"
	"go/adv-example/configs"
	"go/adv-example/db"
	"go/adv-example/internal/handler"
	"log"
	"net/http"
)

func main() {
	conf := configs.NewConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()
	handler.NewHandler(router, conf.Smtp)

	server := http.Server{
		Addr:    conf.Port,
		Handler: router,
	}

	fmt.Print("Server start on port 8081")
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}
