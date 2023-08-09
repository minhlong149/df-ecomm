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
	"df-ecomm/pkg/util"
)

type App struct {
	Logger *log.Logger
	Router *gin.Engine
	Config util.Config
}

func (app *App) SetupRouter() {
	handler := handler.Setup(app.Logger, app.Config)

	productsRoute := app.Router.Group("/products")
	productsRoute.Use(middleware.Authenticated(app.Config.SecretKey))
	{
		productsRoute.GET("", handler.GetAllProducts)
		productsRoute.POST("", handler.AddNewProduct)
		productsRoute.PUT("/:id", handler.UpdateProductById)
		productsRoute.DELETE("/:id", handler.RemoveProductById)
	}

	app.Router.POST("/auth", handler.Login)

	cartRoute := app.Router.Group("/cart")
	{
		cartRoute.POST("/add", handler.AddItem)
		cartRoute.DELETE("/remove", handler.RemoveItem)
		cartRoute.POST("/checkout", handler.Checkout)
	}
}

func main() {
	app := &App{
		Logger: log.Default(),
		Router: gin.Default(),
		Config: util.LoadConfig(),
	}

	app.SetupRouter()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.Config.Port),
		Handler: app.Router,
	}

	go func() {
		app.Logger.Printf("Server listening on port %s\n", app.Config.Port)
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
