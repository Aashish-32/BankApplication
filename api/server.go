package api

import (
	db "github.com/Aashish-32/bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	router := gin.Default()
	server := &Server{store: store}
	router.POST("/accounts", server.createAccount)

	server.router = router

	return server

}

func (server *Server) Start(address string) error {
	return (server.router.Run(address))

}
