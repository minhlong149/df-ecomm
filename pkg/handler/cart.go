package handler

import (
	"github.com/gin-gonic/gin"

	"df-ecomm/pkg/model"
)

func (h *Handler) AddItem(c *gin.Context) {
	var newItem model.ItemToAdd
	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if newItem.Quantity < 0 {
		c.JSON(400, gin.H{"error": "Quantity must be greater than 0"})
		return
	}

	cartItem, err := h.Cart.Add(newItem)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, cartItem)
}

func (h *Handler) RemoveItem(c *gin.Context) {
	var item model.ItemToRemove
	if err := c.BindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.Cart.Remove(item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

func (h *Handler) Checkout(c *gin.Context) {
	receipt := h.Cart.Checkout()
	c.JSON(200, receipt)
}
