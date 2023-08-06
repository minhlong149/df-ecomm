package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"

	"df-ecomm/pkg/handler"
	"df-ecomm/pkg/service"
)

type App struct {
	Logger *log.Logger
	Router *gin.Engine
}

func (app *App) ProductRouter() {
	productHandler := handler.ProductHandler{
		ProductService: &service.ProductService{},
	}

	productsRoute := app.Router.Group("/products")
	{
		productsRoute.GET("/", productHandler.GetAllProducts)
		productsRoute.POST("/", productHandler.AddNewProduct)
		productsRoute.PUT("/:id", productHandler.UpdateProductById)
		productsRoute.DELETE("/:id", productHandler.RemoveProductById)
	}
}

func main() {
	app := &App{
		Logger: log.Default(),
		Router: gin.Default(),
	}

	app.ProductRouter()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: app.Router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			app.Logger.Fatalf("Listen: %s", err)
		}
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	if err := srv.Shutdown(context.Background()); err != nil {
		app.Logger.Fatalf("Server shutdown: %s\n", err)
	}
}
