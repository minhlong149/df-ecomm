package handler

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"df-ecomm/pkg/model"
)

type ProductHandler struct {
	ProductService model.ProductService
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products := h.ProductService.GetAll()
	c.JSON(200, products)
}

func (h *ProductHandler) AddNewProduct(c *gin.Context) {
	var newProduct model.NewProduct
	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if newProduct.Price < 0 {
		c.JSON(400, gin.H{"error": "Price must be greater than 0"})
		return
	}

	createdProduct := h.ProductService.Create(newProduct)
	c.JSON(201, createdProduct)
}

func (h *ProductHandler) UpdateProductById(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid product id"})
		return
	}

	var updateProduct model.Product
	if err := c.BindJSON(&updateProduct); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if updateProduct.Price < 0 {
		c.JSON(400, gin.H{"error": "Price must be greater than 0"})
		return
	}

	updatedProduct, err := h.ProductService.UpdateById(productId, updateProduct)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, updatedProduct)
}

func (h *ProductHandler) RemoveProductById(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid product id"})
		return
	}

	if err := h.ProductService.DeleteById(productId); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.Status(204)
}
