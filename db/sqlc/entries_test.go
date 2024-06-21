package db

import (
	"context"
	"github.com/musstafayilmazz/Golang/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateRandomEntry(t *testing.T) Entry {
	account, err := testQueries.GetRandomAccount(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	arg := CreateEntryParams{
		Amount:    util.RandomMoney(),
		AccountID: account.ID,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.CreatedAt)
	require.NotZero(t, entry.ID)

	return entry
}

func TestCreateEntry(t *testing.T) {
	CreateRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := CreateRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.CreatedAt, entry2.CreatedAt)

}

func TestUpdateEntry(t *testing.T) {
	entry1 := CreateRandomEntry(t)

	arg := UpdateEntryParams{
		ID:        entry1.ID,
		AccountID: entry1.AccountID,
		Amount:    entry1.Amount,
	}
	entry2, err := testQueries.UpdateEntry(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomEntry(t)
	}

	args := ListEntryParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntry(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
