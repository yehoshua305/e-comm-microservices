package db

import (
	"context"
	// "log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yehoshua305/e-comm-microservices/src/util"
)

func createRandomUser(t *testing.T) User {

	hashedPassword, err := util.HashPassword(util.RandomString(8))
	require.NoError(t, err)

	randomUser := User{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		Email:          util.RandomEmail(),
		Phone:          util.RandomPhone(),
		FullName:       util.RandomFullName(),
		Address:        util.RandomAddress(),
		CreatedAt:      time.Now(),
	}

	// Write User to the Dynamodb Table
	User, err := testDB.CreateUser(context.Background(), randomUser)
	require.NoError(t, err)
	require.NotEmpty(t, User)

	return randomUser
}

func TestCreateUser(t *testing.T) {
	_ = createRandomUser(t)

	// require.Equal(t, randomUser.Username, User.Username)
	// require.Equal(t, randomUser.Email, User.Email)
	// require.Equal(t, randomUser.Phone, User.Phone)
	// require.Equal(t, randomUser.FullName, User.FullName)
	// require.Equal(t, randomUser.Address, User.Address)

}

func TestGetUser(t *testing.T) {
	randomUser := createRandomUser(t)

	User2, err := testDB.GetUser(context.Background(), randomUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, User2)

	require.Equal(t, randomUser.Username, User2.Username)
	require.Equal(t, randomUser.Email, User2.Email)
	require.Equal(t, randomUser.Phone, User2.Phone)
	require.Equal(t, randomUser.FullName, User2.FullName)
	require.Equal(t, randomUser.Address, User2.Address)
	require.WithinDuration(t, randomUser.CreatedAt, User2.CreatedAt, time.Second)
	// Compare timestamps without considering the monotonic clock reading.
	// require.True(t, randomUser.CreatedAt.Equal(User2.CreatedAt), "CreatedAt timestamps should be equal")
}

func TestDeleteUser(t *testing.T) {
	randomUser := createRandomUser(t)

	_, err := testDB.DeleteUser(context.Background(), randomUser.Username)
	require.NoError(t, err)

	User2, err := testDB.GetUser(context.Background(), randomUser.Username)
	require.Error(t, err)
	require.Empty(t, User2)
}
