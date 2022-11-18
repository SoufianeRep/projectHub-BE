package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
}

// Global variable to be used elswhere
var Router *gin.Engine = gin.Default()

// NewServer creates and returns an instance of a server
func NewServer(db *gorm.DB) *Server {
	server := &Server{router: Router}
	Router.MaxMultipartMemory = 8 << 20

	Router.POST("/upload", handleUpload)

	Router.POST("/users/create", handleCreateUser)

	Router.POST("/teams", handleCreateTeam)
	Router.POST("/teams/:id/members", handleAddMemberToTeam)
	Router.POST("/login", handleLogin)

	return server
}

// Start runs a gin default server on the given address
func (server *Server) Start(serverAdr string) error {
	return server.router.Run(serverAdr)
}

// func errorResponse(err error) gin.H {
// 	return gin.H{"message": err.Error()}
// }
