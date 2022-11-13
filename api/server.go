package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

// Global variable to be used elswhere
var Router *gin.Engine = gin.Default()

// NewServer creates and returns an instance of a server
func NewServer() *Server {
	server := &Server{router: Router}
	Router.MaxMultipartMemory = 8 << 20

	Router.POST("/upload", handleUpload)

	return server
}

// Start runs a gin default server on the given address
func (server *Server) Start(serverAdr string) error {
	return server.router.Run(serverAdr)
}