package util

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v5"
)

// Claims is a struct that contains the claims of a token
// RegisteredClaims - iss, sub, aud, exp, nbf, iat, and jti
// Public claims - name and email.
// Private claims - user_id and role.
type Claims struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	jwt.RegisteredClaims 
}

// NewClaims create a new claim for a specific user
func NewClaims(username string, duration time.Duration) (*Claims, error) {
	if username == "" {
		return nil, errors.New("username is required")
	}

	if duration == 0 {
		return nil, errors.New("duration is required")
	}

	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}	

	claims := &Claims{
		ID: tokenID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	return claims, nil
}

// Check validity of token claims
func (claims *Claims) ClaimsValid() error {
	expTime, err := claims.GetExpirationTime()
	if err != nil {
		return err
	}
	if time.Now().After(expTime.Time) {
		return jwt.ErrTokenExpired
	}
	return nil
}