package users

import (
	"fmt"

	"github.com/gin-gonic/gin" 
	"github.com/yehoshua305/e-comm-microservices/src/db"
	"github.com/yehoshua305/e-comm-microservices/src/token"
	"github.com/yehoshua305/e-comm-microservices/src/util"
)

// Server serves HTTP requests
type Server struct {
	config util.Config
	table  db.Table
	router *gin.Engine
	tokenMaker token.Maker
}

func NewServer(config util.Config, table db.Table) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config: config,
		table: table,
		tokenMaker: tokenMaker,
	}
	server.setupRouter()
	return server, nil
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	// initializes a new instance of the Gin web framework's router
	// with default middleware
	router := gin.Default()

	router.POST("/user", server.createUser)
	router.POST("user/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	authRoutes := router.Group("/").Use(util.AuthMiddleware(server.tokenMaker))
	authRoutes.PUT("/user", server.updateUser)
	authRoutes.GET("/user/:username", server.getUser)
	authRoutes.DELETE("/user/:username", server.deleteUser)

	server.router = router
}