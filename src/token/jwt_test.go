package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
	"github.com/yehoshua305/e-comm-microservices/src/util"
)

func TestNewJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, claims, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, claims)

	claims, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, claims)

	require.NotZero(t, claims.ID)
	require.Equal(t, username, claims.Username)
	claimsIssuedAt, err := claims.GetIssuedAt()
	require.NoError(t, err)
	claimsExpiredAt, err := claims.GetExpirationTime()
	require.NoError(t, err)
	require.WithinDuration(t, issuedAt, claimsIssuedAt.Time, time.Second)
	require.WithinDuration(t, expiredAt, claimsExpiredAt.Time, time.Second)
}

func TestExpiredToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	token, claims, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, claims)

	claims, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.Nil(t, claims)
}

func TestInvalidToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	claims, err := NewClaims(util.RandomOwner(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, claims)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	claims, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, maker.ParseError(err), ErrInvalidToken.Error())
	require.Nil(t, claims)
}