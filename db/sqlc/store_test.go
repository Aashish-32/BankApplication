package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	fmt.Println("<<before", account1.Balance, account2.Balance)

	//run n concurrent transfer transaction

	n := 2
	amount := int64(10)
	errs := make(chan error)

	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {

		txname := fmt.Sprintf("transaction %d", i+1)
		go func() {
			ctx := context.WithValue(context.Background(), txkey, txname)
			result, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}

	//check results

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		transfer := result.Transfer

		require.NotEmpty(t, result)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		//check entries
		FromEntry := result.FromEntry
		require.NotEmpty(t, FromEntry)

		require.Equal(t, account1.ID, FromEntry.AccountID)
		require.Equal(t, -amount, FromEntry.Amount)
		require.NotZero(t, FromEntry.ID)
		require.NotZero(t, FromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), FromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)

		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		//check accounts
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.ID, toAccount.ID)

		//TODO: check accounts balance

		fmt.Println(">>tx:", fromAccount.Balance, toAccount.Balance)

		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff2 > 0)
		require.True(t, diff1%amount == 0)

		//check final updated balances

		updatedaccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
		require.NoError(t, err)

		updatedaccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
		require.NoError(t, err)

		fmt.Println("<<after", updatedaccount1.Balance, updatedaccount2.Balance)

		// require.Equal(t, account1.Balance-int64(n)*amount, updatedaccount1.Balance)
		// require.Equal(t, account2.Balance+int64(n)*amount, updatedaccount2.Balance)

	}

}
