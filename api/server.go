package api

import (
	"fmt"
	"os"

	"github.com/Aashish-32/bank/token"

	db "github.com/Aashish-32/bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store      *db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(store *db.Store) (*Server, error) {
	token_key := os.Getenv("token_symmetric_key")
	// token_duration := os.Getenv("Access_token_duration")

	tokenMaker, err := token.NewPasetoMaker(token_key)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %v", err)
	}

	// new_token_duration, err := time.ParseDuration(token_duration)
	// if err != nil {
	// 	return nil, fmt.Errorf("unable parse duration: %v", err)
	// }

	// token, err := tokenMaker.CreateToken(token_key, new_token_duration)
	// if err != nil {
	// 	return nil, fmt.Errorf("unable to create token: %v", err)
	// }
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.POST("/transfers", server.createTransfer)
	router.POST("/users", server.createUser)
	router.GET("/users/:username", server.getUser)

	server.router = router

	return server, nil

}

func (server *Server) Start(address string) error {
	return (server.router.Run(address))

}
