package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	
	password := RandomString(8)
	hashedPassword1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	err = CheckPassword(hashedPassword1, password)
	require.NoError(t, err)

	badActorPassword := RandomString(8)
	err = CheckPassword(hashedPassword1, badActorPassword)
	require.Error(t, err, bcrypt.ErrMismatchedHashAndPassword)

	hashedPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}

