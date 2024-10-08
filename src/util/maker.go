package util

import "time"

type Maker interface {
	// CreateToken creates a new token for a specific username
	CreateToken(username string, duration time.Duration) (string, *Claims, error)

	// VerifyToken checks if the token is valid
	VerifyToken(token string) (*Claims, error) 

	// ParseError
	ParseError(err error) error
}