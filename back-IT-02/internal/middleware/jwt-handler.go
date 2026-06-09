package middleware

import (
	"net/http"
	"strings"

	"github.com/Rafiana219/interview-question-02/back-IT-02/internal/utils"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		auth := c.GetHeader("Authorization")

		if auth == "" {

			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "missing token",
			})

			c.Abort()
			return
		}

		token := strings.TrimPrefix(auth, "Bearer ")

		userID, err := utils.ValidateToken(token)

		if err != nil {

			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})

			c.Abort()
			return
		}

		c.Set("userId", userID)

		c.Next()
	}
}
