package middlewares

import (
	"fmt"
	"net/http"

	"github.com/briannbig/habitify/internal/util"
	"github.com/gin-gonic/gin"
)

const (
	strAuthorizationHeader = "Authorization"
	strBearerSchema        = "Bearer "
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeaderStr := ctx.GetHeader(strAuthorizationHeader)
		tokenStr := authHeaderStr[len(strBearerSchema):]
		if tokenStr == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unaouthorized"})
			return
		}
		err := util.VerifyToken(tokenStr)

		if err != nil {
			errStr := fmt.Sprintf("error parsing token --- %s", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errStr})
			return
		}

	}
}
