package db

import (
	"context"
	"github.com/musstafayilmazz/Golang/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomTransfer(t *testing.T) Transfer {
	account1, err := testQueries.GetRandomAccount(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	account2, err := testQueries.GetRandomAccount(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}
	require.NotEmpty(t, transfer)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.NotZero(t, transfer.CreatedAt)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	CreateRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := CreateRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, time.Now(), transfer1.CreatedAt, time.Second)

}

func TestGetAllTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomTransfer(t)
	}
	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)

	}
}
