package middleware

import (
	"strings"

	"df-ecomm/pkg/model"
	"df-ecomm/pkg/util"

	"github.com/gin-gonic/gin"
)

func Authenticated(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.Split(c.GetHeader("Authorization"), " ")
		if len(authHeader) != 2 {
			c.AbortWithStatusJSON(401, gin.H{"error": model.ErrInvalidAuthHeader.Error()})
			return
		}

		switch authHeader[0] {
		case "Bearer":
			claims, ok := util.VerifyToken(authHeader[1], secretKey)
			if !ok {
				c.AbortWithStatusJSON(401, gin.H{"error": model.ErrInvalidToken.Error()})
				return
			}

			if claims["role"] != "admin" {
				c.AbortWithStatusJSON(403, gin.H{"error": model.ErrUnauthorized.Error()})
				return
			}
		default:
			c.AbortWithStatusJSON(401, gin.H{"error": model.ErrInvalidAuthHeader.Error()})
			return
		}

		c.Next()
	}
}
