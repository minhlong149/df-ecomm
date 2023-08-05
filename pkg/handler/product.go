package handler

import (
	"github.com/gin-gonic/gin"
	"df-ecomm/pkg/model"
)

type ProductHandler struct {
	ProductService model.ProductService
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	products := h.ProductService.GetAll()
	c.JSON(200, products)
}
