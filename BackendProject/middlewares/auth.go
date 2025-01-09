package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/backend/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized.",
		})
		return
	}

	userId, verifyErr := utils.VerifyToken(token)

	if verifyErr != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized.",
			"error":   verifyErr,
		})

		return
	}

	context.Set("userId", userId)
	context.Next()
}
