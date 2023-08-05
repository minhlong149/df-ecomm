package main

import (
	"github.com/gin-gonic/gin"
	"df-ecomm/pkg/handler"
	"df-ecomm/pkg/service"
)

func main() {
	r := gin.Default()

	productHandler := handler.ProductHandler{
		ProductService: &service.ProductService{},
	}

	r.GET("/products", productHandler.GetAll)

	r.Run()
}
