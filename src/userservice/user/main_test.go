package user

import (
	"testing"
	"time"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/yehoshua305/e-comm-microservices/src/db"
	"github.com/yehoshua305/e-comm-microservices/src/util"
)

func newTestServer(t *testing.T, table db.Table) *Server {
	config := util.Config{
		TokenSymmetricKey: util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, table)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
