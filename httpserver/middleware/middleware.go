package middleware

import (
	"net/http"

	"github.com/deevarindu/final-project-2/helper/jwt"
	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {
	verifyToken, err := jwt.VerifyToken(ctx)
	_ = verifyToken
	// return
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthenticated",
			"message": err.Error(),
		})
		return
	}
	ctx.Set("userData", verifyToken)
	ctx.Next()
}
