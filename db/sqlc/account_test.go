package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Aashish-32/bank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NotEmpty(t, account)
	require.NoError(t, err)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}
func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)

}

func TestGetAccount(t *testing.T) {
	acc := CreateRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(), acc.ID)
	require.NotEmpty(t, account)
	require.NoError(t, err)
	require.Equal(t, acc.Owner, account.Owner)
	require.Equal(t, acc.Balance, account.Balance)
	require.Equal(t, acc.Currency, account.Currency)
	require.WithinDuration(t, acc.CreatedAt, account.CreatedAt, time.Second)

}
func TestUpdateAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	args := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}
	acc2, err := testQueries.UpdateAccount(context.Background(), args)
	require.Equal(t, args.Balance, acc2.Balance)
	require.NotEqual(t, account1.Balance, acc2.Balance)
	require.Equal(t, account1.ID, acc2.ID)
	require.Equal(t, account1.Currency, acc2.Currency)
	require.Equal(t, account1.Owner, acc2.Owner)

	require.NoError(t, err)
	require.NotEmpty(t, acc2)
	require.WithinDuration(t, account1.CreatedAt, acc2.CreatedAt, time.Second)
}
func TestDeleteAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)

}
func TestListAccount(t *testing.T) {
	var lastac Account
	for i := 0; i < 6; i++ {
		lastac = CreateRandomAccount(t)
	}
	arg := ListAccountsParams{
		Owner:  lastac.Owner,
		Limit:  5,
		Offset: 0,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastac.Owner, account.Owner)

	}
}
