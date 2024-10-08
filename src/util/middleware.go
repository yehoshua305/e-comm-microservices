package util

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeaderKey = "authorization"
	AuthorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

// authMiddleware creates a gin middleware for authorization.
// It takes a tokenMaker as a parameter, which is responsible for
// creating and verifying access tokens.
// The middleware checks the authorization header in the request,
// verifies the access token, and sets the authorization payload
// in the context for further processing.
func AuthMiddleware(tokenMaker Maker) gin.HandlerFunc {
	return func (ctx *gin.Context) {
		// get authorization header
		authorizationHeader := ctx.GetHeader(AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header required")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse(err))
			return
		}

		// split authorization and get fields
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse(err))
			return
		}

		// check authorization header type
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != AuthorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse(err))
			return
		}
			
		// verify authorization token
		accessToken := fields[1]
		claims, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse(err))
			return
		}

		ctx.Set(AuthorizationPayloadKey, claims)

		// The Next method is responsible for advancing the execution to 
		// the next middleware in the chain.
		ctx.Next()
	}
}