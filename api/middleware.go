package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Aashish-32/bank/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationPayloadKey = "authorization_key"
	authorizationTypeBearer = "bearer"
)

var authorizationTypes = []string{"bearer"}

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
		authorizationType := strings.ToLower(fields[0])
		for i := range authorizationTypes {
			if authorizationType != authorizationTypes[i] {
				err := fmt.Errorf("unsupported authorization type %s", authorizationType)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
				return

			}
		}
		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()

	}
}
