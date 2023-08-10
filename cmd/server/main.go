package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"df-ecomm/pkg/handler"
	"df-ecomm/pkg/middleware"
	"df-ecomm/pkg/util"
)

type App struct {
	Db     *gorm.DB
	Router *gin.Engine
	Logger util.Logger
	Config util.Config
}

func (app *App) SetupRouter() {
	handler := handler.Setup(app.Logger, app.Config, app.Db)

	app.Router.Use(middleware.ErrorHandler(app.Logger))

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
		Router: gin.Default(),
		Logger: util.SetupLogger(),
		Config: util.LoadConfig(),
	}

	app.Db = util.SetupDB(app.Config)
	app.SetupRouter()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.Config.Port),
		Handler: app.Router,
	}

	go func() {
		app.Logger.Info("Server listening on port %s\n", app.Config.Port)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			app.Logger.Error("Listen: %s", err)
		}
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	if err := srv.Shutdown(context.Background()); err != nil {
		app.Logger.Error("Server shutdown: %s\n", err)
	}
}
