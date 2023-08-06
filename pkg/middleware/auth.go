package middleware

import (
	"strings"

	"df-ecomm/pkg/util"

	"github.com/gin-gonic/gin"
)

func Authenticated(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.Split(c.GetHeader("Authorization"), " ")
		if len(authHeader) != 2 {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid Authorization header"})
			return
		}

		switch authHeader[0] {
		case "Bearer":
			_, ok := util.VerifyToken(authHeader[1], secretKey)
			if !ok {
				c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
				return
			}
		default:
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid Authorization header"})
		}

		c.Next()
	}
}
