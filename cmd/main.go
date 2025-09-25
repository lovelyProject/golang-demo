package main

import (
	"go/adv-example/configs"
	"go/adv-example/db"
	"go/adv-example/internal/link"
	"go/adv-example/internal/product"
	"go/adv-example/pkg/middleware"
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
	// repository
	database := db.NewDb(conf)
	linkRepo := link.NewLinkRepository(database)
	linkService := link.NewService(linkRepo)

	productRepo := product.NewProductRepo(database)
	productService := product.NewProductService(productRepo)
	// handler
	productHandler := product.NewProductHandler(productService)
	productHandler.RegisterRoutes(router)
	linkHandler := link.NewLinkHandler(link.LinkHandlerDeps{
		LinkService: linkService,
	})
	linkHandler.RegisterRoutes(router)

	middlewares := middleware.Chain(
		middleware.IsAuthenticated,
		middleware.Log,
	)
	server := http.Server{
		Addr:    conf.Port,
		Handler: middlewares(router),
	}

	log.Println("server started")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}

}
