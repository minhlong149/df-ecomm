package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"df-ecomm/pkg/model"
)

func (h *Handler) GetAllProducts(c *gin.Context) {
	products, err := h.Product.GetAll()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{"products": products})
}

func (h *Handler) AddNewProduct(c *gin.Context) {
	var newProduct model.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": model.ErrInvalidProduct.Error()})
		return
	}

	createdProduct, err := h.Product.Create(newProduct)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, gin.H{"product": createdProduct})
}

func (h *Handler) UpdateProductById(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": model.ErrInvalidProductId.Error()})
		return
	}

	var productToUpdate model.Product
	if err := c.ShouldBindJSON(&productToUpdate); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": model.ErrInvalidProduct.Error()})
		return
	}

	updatedProduct, err := h.Product.UpdateById(productId, productToUpdate)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{"product": updatedProduct})
}

func (h *Handler) RemoveProductById(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": model.ErrInvalidProductId.Error()})
		return
	}

	if err := h.Product.DeleteById(productId); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}
