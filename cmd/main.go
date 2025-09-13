package main

import (
	"go/adv-example/configs"
	"go/adv-example/db"
	handlr "go/adv-example/internal/handler"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	conf := configs.NewConfig()
	router := http.NewServeMux()
	handlr.NewHandler(router)
	db.NewDb(conf)
	server := http.Server{
		Addr:    conf.Port,
		Handler: router,
	}

	log.Println("server started")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}

}
