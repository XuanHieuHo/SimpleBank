package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/techschool/simplebank/db/sqlc"
)

//Server serves HTTP requests for our banking service
type Server struct {
	store db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// add routes to router
	// accounts
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id",server.getAccount)
	router.GET("/accounts",server.listAccount)
	router.PUT("/accounts/:id",server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	//transfers
	router.POST("/transfers", server.createTransfer)
	router.GET("/transfers/:id", server.getTransfer)
	router.GET("/transfers/list/from/:from_account_id/", server.listTransferFrom)
	router.GET("/transfers/list/to/:to_account_id/", server.listTransferTo)
	server.router = router
	return server
}

// Start runs thes HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
} 

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}