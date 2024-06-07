package api

import (
	"database/sql"
	"net/http"
	"time"

	"os"

	"log"

	db "github.com/Aashish-32/bank/db/sqlc"
	"github.com/Aashish-32/bank/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=7"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}
type UserResponse struct {
	Username            string
	FullName            string
	Email               string
	Password_changed_at time.Time
	Created_at          time.Time
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashed_password, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashed_password,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	account, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqerr, ok := err.(*pq.Error); ok {
			ctx.JSON(http.StatusForbidden, gin.H{"Error": pqerr.Message, "severity": pqerr.Severity})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, account)

}

type getUserparams struct {
	Username string `uri:"username" binding:"required,alphanum"`
}

func Newuserresponse(user db.User) UserResponse {
	return UserResponse{
		Username:            user.Username,
		FullName:            user.FullName,
		Email:               user.Email,
		Password_changed_at: user.PasswordChangedAt,
		Created_at:          user.CreatedAt,
	}

}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserparams
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	account, err := server.store.GetUser(ctx, req.Username)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := Newuserresponse(account)
	ctx.JSON(http.StatusOK, response)

}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=7"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user_response"`
}

func (server *Server) login(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return

		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	newuser := Newuserresponse(user)
	checkPasswordErr := util.CheckPassword(req.Password, user.HashedPassword)
	if checkPasswordErr != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": checkPasswordErr.Error()})
		return
	}

	token_duration := os.Getenv("Access_token_duration")

	new_token_duration, err := time.ParseDuration(token_duration)
	if err != nil {
		log.Println("unable parse duration: %v", err)
	}

	accessToken, err := server.tokenMaker.CreateToken(req.Username, new_token_duration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := loginUserResponse{
		AccessToken: accessToken,
		User:        newuser,
	}

	ctx.JSON(http.StatusAccepted, response)

}
