package users

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yehoshua305/golang-simple-api/token"
	"github.com/yehoshua305/golang-simple-api/util"
)

const (
	authorizationHeaderKey = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// authMiddleware creates a gin middleware for authorization.
// It takes a tokenMaker as a parameter, which is responsible for
// creating and verifying access tokens.
// The middleware checks the authorization header in the request,
// verifies the access token, and sets the authorization payload
// in the context for further processing.
func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func (ctx *gin.Context) {
		// get authorization header
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header required")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrorResponse(err))
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
		if authorizationType != authorizationTypeBearer {
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

		ctx.Set(authorizationPayloadKey, claims)

		// The Next method is responsible for advancing the execution to 
		// the next middleware in the chain.
		ctx.Next()
	}
}