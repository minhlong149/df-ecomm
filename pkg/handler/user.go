package handler

import (
	"github.com/gin-gonic/gin"

	"df-ecomm/pkg/model"
)

type UserHandler struct {
	UserService model.UserService
}

func (h *UserHandler) Login(c *gin.Context) {
	var account model.Account
	if err := c.BindJSON(&account); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserService.Login(account)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}
