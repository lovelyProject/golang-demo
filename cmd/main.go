package main

import (
	"go/adv-example/configs"
	"go/adv-example/db"
	"log"
	"net/http"
)

func main() {
	conf := configs.NewConfig()
	router := http.NewServeMux()
	db.NewDb(conf)
	server := http.Server{
		Addr:    conf.Port,
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
