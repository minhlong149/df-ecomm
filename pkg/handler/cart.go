package handler

import (
	"github.com/gin-gonic/gin"

	"df-ecomm/pkg/model"
)

type ItemToAdd struct {
	ProductId uint `json:"id" binding:"required"`
	Quantity  uint `json:"quantity" binding:"required"`
}

type ItemToRemove struct {
	ProductId uint `json:"id" binding:"required"`
}

type Receipt struct {
	Items []model.CartItem `json:"cart"`
	Total uint             `json:"total"`
}

func (h *Handler) AddItem(c *gin.Context) {
	var newItem ItemToAdd
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": model.ErrInvalidCart.Error()})
		return
	}

	cartItem, err := h.Cart.Add(newItem.ProductId, newItem.Quantity)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, cartItem)
}

func (h *Handler) RemoveItem(c *gin.Context) {
	var item ItemToRemove
	if err := c.ShouldBindJSON(&item); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": model.ErrInvalidCart.Error()})
		return
	}

	if err := h.Cart.Remove(item.ProductId); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}

func (h *Handler) Checkout(c *gin.Context) {
	cartItems, err := h.Cart.Checkout()
	if err != nil {
		c.Error(err)
		return
	}

	receipt := Receipt{Items: cartItems}
	for _, item := range cartItems {
		receipt.Total += item.Quantity * item.Product.Price
	}

	c.JSON(200, receipt)
}
