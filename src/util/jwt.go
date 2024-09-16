package util

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

// Create JWTMaker
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be atleast %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}


// Create Token for a username
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, *Claims, error) {
	claims, err := NewClaims(username, duration)
	if err != nil {
		return "", &Claims{}, err
	}

	// create token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signed token
	token, err := jwtToken.SignedString([]byte(maker.secretKey))

	return token, claims, err
}

// VerifyToken checks the validity of a token
func (maker *JWTMaker) VerifyToken(token string) (*Claims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(maker.secretKey), nil
	}
		
	jwtToken, err := jwt.ParseWithClaims(token, &Claims{}, keyFunc)
	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid token")
	}
	return claims, nil


}

// Parse Error
func (maker *JWTMaker) ParseError(err error) error {
	if err == nil {
		return nil
	}
	splitErr := (strings.Split(err.Error(), ":"))
	splitErr1 := splitErr[len(splitErr)-1]
	return errors.New(strings.TrimSpace(splitErr1))
}