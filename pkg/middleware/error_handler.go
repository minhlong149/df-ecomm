package middleware

import (
	"github.com/gin-gonic/gin"

	"df-ecomm/pkg/model"
	"df-ecomm/pkg/util"
)

func ErrorHandler(logger util.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			switch err.Err {

			case model.ErrWrongCredentials:
				c.JSON(401, gin.H{"error": err.Error()})

			case model.ErrProductNotFound:
				c.JSON(404, gin.H{"error": err.Error()})

			default:
				logger.Error(err)
				c.JSON(500, gin.H{"error": "Something went wrong"})
				return
			}
		}
	}
}
