package handler

import (
	"github.com/gin-gonic/gin"

	"df-ecomm/pkg/model"
)

func (h *Handler) Login(c *gin.Context) {
	var account model.Account
	if err := c.BindJSON(&account); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := h.User.Login(account, h.Config.SecretKey)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}
