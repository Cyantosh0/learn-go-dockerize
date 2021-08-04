package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/cyantosh0/dockerize-go-app/config"
	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context) {
	header := c.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(header, "Bearer", "", 1))

	_, err := config.FB.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	c.Next()
}
