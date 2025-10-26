package main

import (
	"go/adv-example/configs"
	"go/adv-example/db"
	"go/adv-example/internal/auth"
	"go/adv-example/internal/link"
	"go/adv-example/internal/order"
	"go/adv-example/internal/product"
	"go/adv-example/internal/stat"
	"go/adv-example/pkg/event"
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

	eventBus := event.NewEventBus()

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
		Config:      conf,
	}, conf)

	//auth
	authRepo := auth.NewAuthRepo(database)
	authService := auth.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(authService, conf)
	authHandler.RegisterRoutes(router)

	linkHandler.RegisterRoutes(router)

	//stat
	statRepo := stat.NewStatRepository(database)
	statService := stat.NewStatService(statRepo, eventBus)
	statHanlder := stat.NewStatHandler(statService, conf)
	statHanlder.RegisterRoutes(router)

	// order
	orderRepo := order.NewOrderRepository(database)
	orderService := order.NewService(orderRepo)
	orderHandler := order.NewOrderHandler(orderService, conf)
	orderHandler.RegisterRoutes(router)

	middlewares := middleware.Chain(
		// middleware.IsAuthenticated,
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
