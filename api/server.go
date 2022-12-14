package api

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/SoufianeRep/tscit/token"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var TD time.Duration

func init() {
	// Convert token duration value from env variable
	td, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_DURATION"))
	if err != nil {
		log.Fatal("unable to convert token variable duration", err)
	}
	TD = time.Duration(td)
}

// Server serves HTTP requests for the app
type Server struct {
	router     *gin.Engine
	tokenMaker token.Maker
	db         *gorm.DB
}

// NewServer creates and returns an instance of a server
func NewServer(db *gorm.DB) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(os.Getenv("TOKEN_SYMMETRIC_KEY"))
	if err != nil {
		return nil, fmt.Errorf("cant create a token: %v", err)
	}

	server := &Server{
		tokenMaker: tokenMaker,
		db:         db,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20

	router.POST("/users/signup", handleCreateUser)
	router.POST("/users/login", server.handleLogin)

	authRoutes := router.Group("/").Use(authMidldeware(server.tokenMaker))

	authRoutes.GET("/users/:id", handleGetUser)                  // Get the user information
	authRoutes.GET("/users/:id/teams", handleGetUserTeams)       // Get all the the teams the user is part of
	authRoutes.POST("/teams", handleCreateTeam)                  // Create a new team
	authRoutes.GET("/teams/:id", handleGetTeam)                  // Get team information with all team members and projects optionally
	authRoutes.POST("/teams/:id/members", handleAddMemberToTeam) // Add a new member to a team manually
	router.POST("/teams/:id/projects", handleCreateProject)
	authRoutes.GET("/projects/:id", handleGetProject)
	authRoutes.POST("/upload", handleUpload)

	server.router = router
}

// Start runs a gin default server on the given address
func (server *Server) Start(serverAdr string) error {
	return server.router.Run(serverAdr)
}
