package db

import (
	"context"
	"testing"

	"github.com/Frank2006x/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	errs:=make(chan error)
	results:=make(chan TransferTxResult)

	n:=5
	

	for i := 0; i < n; i++ {
		go func() {
			arg := CreateTransferParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        util.RandomMoney(),
			}
			result, err := store.TransferTx(context.Background(), arg)
			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
		result := <-results
		require.NotEmpty(t, result)

		require.Equal(t, account1.ID, result.Transfer.FromAccountID)
		require.Equal(t, account2.ID, result.Transfer.ToAccountID)
		require.Equal(t, -result.Transfer.Amount, result.FromEntry.Amount)
		require.Equal(t, result.Transfer.Amount, result.ToEntry.Amount)
		require.NotZero(t, result.Transfer.ID)
		require.NotZero(t, result.Transfer.CreatedAt)


		//TODO: check accounts' balance
	}

}




