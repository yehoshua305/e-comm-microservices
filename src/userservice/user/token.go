package user

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yehoshua305/e-comm-microservices/src/util"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type renewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (server *Server) renewAccessToken(ctx *gin.Context) {
	var req renewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	// Verify Refresh Token
	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, util.ErrorResponse(err))
		return
	}

	// Get Session
	session, err := server.table.GetSession(ctx, refreshPayload.Username, refreshPayload.ID.String())
	if err != nil {
		ctx.JSON(http.StatusNotFound, util.ErrorResponse(err))
		return
	}

	log.Println("SessionID: ", session.ID)

	// Check if Session is blocked
	if session.IsBlocked {
		ctx.JSON(http.StatusUnauthorized, util.ErrorResponse(fmt.Errorf("blocked session")))
		return
	}

	// Check username
	if session.Username != refreshPayload.Username {
		ctx.JSON(http.StatusUnauthorized, util.ErrorResponse(fmt.Errorf("incorrect session user")))
		return
	}

	// Check refreshToken
	if session.RefreshToken != req.RefreshToken {
		ctx.JSON(http.StatusUnauthorized, util.ErrorResponse(fmt.Errorf("mismatch session token")))
		return
	}

	// Check expiration time
	if time.Now().After(session.ExpiresAt) {
		ctx.JSON(http.StatusUnauthorized, util.ErrorResponse(fmt.Errorf("expired session")))
		return
	}

	// Create new Access Token
	accessToken, accessPayload, err := server.tokenMaker.CreateToken(refreshPayload.Username, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	resp := renewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiresAt.Time,
	}

	ctx.JSON(http.StatusOK, resp)
}