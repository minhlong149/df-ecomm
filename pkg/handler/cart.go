package handler

import (
	"github.com/gin-gonic/gin"

	"df-ecomm/pkg/model"
)

type CartHandler struct {
	CartService model.CartService
}

func (h *CartHandler) AddItem(c *gin.Context) {
	var newItem model.ItemToAdd
	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if newItem.Quantity < 0 {
		c.JSON(400, gin.H{"error": "Quantity must be greater than 0"})
		return
	}

	cartItem, err := h.CartService.AddItem(newItem)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, cartItem)
}

func (h *CartHandler) RemoveItem(c *gin.Context) {
	var item model.ItemToRemove
	if err := c.BindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.CartService.RemoveItem(item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

func (h *CartHandler) Checkout(c *gin.Context) {
	receipt := h.CartService.Checkout()
	c.JSON(200, receipt)
}
