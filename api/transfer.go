package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/Aashish-32/bank/db/sqlc"
	"github.com/Aashish-32/bank/token"
	"github.com/gin-gonic/gin"
)

type TransferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,oneof=USD EUR GBP"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req TransferRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fromAccount, valid := server.validAccount(ctx, req.FromAccountID, req.Currency, req.Amount)
	if !valid {
		return
	}
	authpayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authpayload.Username != fromAccount.Owner {
		err := fmt.Errorf("from account %v doesnot belong to the authenticated user", fromAccount.Owner)

		ctx.JSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
		return
	}

	_, valid = server.validAccount(ctx, req.ToAccountID, req.Currency, req.Amount)
	if !valid {
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	transfer, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, transfer)

}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string, amount int64) (db.Account, bool) {
	acc, err := server.store.GetAccount(ctx, accountID)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return acc, false
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return acc, false
	}

	if acc.Currency != currency {
		err := fmt.Errorf(" account %v  currency mismatch: %v vs %v ", acc.ID, acc.Currency, currency)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return acc, false

	}
	if acc.Balance < amount {
		err := fmt.Errorf("insufficient balance in account %v: %v < %v", acc.ID, acc.Balance, amount)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return acc, false
	}
	return acc, true

}
