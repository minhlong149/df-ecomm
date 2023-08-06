package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"

	"df-ecomm/pkg/handler"
	"df-ecomm/pkg/middleware"
	"df-ecomm/pkg/service"
	"df-ecomm/pkg/util"
)

type App struct {
	Logger *log.Logger
	Router *gin.Engine
	Config util.Config
}

func (app *App) ProductRouter() {
	productHandler := handler.ProductHandler{
		ProductService: &service.ProductService{},
	}

	productsRoute := app.Router.Group("/products")
	productsRoute.Use(middleware.Authenticated(app.Config.SecretKey))
	{
		productsRoute.GET("/", productHandler.GetAllProducts)
		productsRoute.POST("/", productHandler.AddNewProduct)
		productsRoute.PUT("/:id", productHandler.UpdateProductById)
		productsRoute.DELETE("/:id", productHandler.RemoveProductById)
	}
}

func (app *App) UserRouter() {
	userHandler := handler.UserHandler{
		UserService: &service.UserService{
			SecretKey: app.Config.SecretKey,
		},
	}
	app.Router.POST("/auth", userHandler.Login)
}

func main() {
	app := &App{
		Logger: log.Default(),
		Router: gin.Default(),
		Config: util.LoadConfig(),
	}

	app.ProductRouter()
	app.UserRouter()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.Config.Port),
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
