package api

import (
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our room service.
type Server struct {
	store  Store
	router *gin.Engine
}

// Router returns the router instance of the server.
func (server *Server) Router() *gin.Engine {
	return server.router
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Change to GET method for fetching room data
	router.GET("/:room_id", server.getRoomData)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
