package handler

import (
	"github.com/gin-gonic/gin"

	"df-ecomm/pkg/model"
)

func (h *Handler) Login(c *gin.Context) {
	var account model.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": model.ErrInvalidCredentials.Error()})
		return
	}

	token, err := h.User.Login(account, h.Config.SecretKey)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{"token": token})
}
