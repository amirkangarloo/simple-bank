package db

import (
	"context"
	"simple-bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type CreateRandomTransferParams struct {
	fromAccount Account
	toAccount   Account
}

func createRandomTransfer(t *testing.T, data CreateRandomTransferParams) Transfer {
	arg := CreateTransferParams{
		Amount:        util.RandomMoney(),
		ToAccountID:   data.toAccount.ID,
		FromAccountID: data.fromAccount.ID,
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.Amount, arg.Amount)
	require.Equal(t, transfer.ToAccountID, arg.ToAccountID)
	require.Equal(t, transfer.FromAccountID, arg.FromAccountID)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	arg := CreateRandomTransferParams{
		toAccount:   createRandomAccount(t),
		fromAccount: createRandomAccount(t),
	}

	createRandomTransfer(t, arg)
}

func TestGetTransfer(t *testing.T) {
	arg := CreateRandomTransferParams{
		toAccount:   createRandomAccount(t),
		fromAccount: createRandomAccount(t),
	}
	transfer := createRandomTransfer(t, arg)
	fetchTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchTransfer)

	require.Equal(t, fetchTransfer.ID, transfer.ID)
	require.Equal(t, fetchTransfer.Amount, transfer.Amount)
	require.Equal(t, fetchTransfer.ToAccountID, transfer.ToAccountID)
	require.Equal(t, fetchTransfer.FromAccountID, transfer.FromAccountID)
	require.WithinDuration(t, fetchTransfer.CreatedAt, transfer.CreatedAt, time.Second)
}

func TestListTransfersOfDeposit(t *testing.T) {
	data := CreateRandomTransferParams{
		toAccount:   createRandomAccount(t),
		fromAccount: createRandomAccount(t),
	}
	for i := 0; i < 10; i++ {
		createRandomTransfer(t, data)
	}

	arg := ListTransfersOfDepositParams{ToAccountID: data.toAccount.ID, Limit: 5, Offset: 5}

	transfers, err := testQueries.ListTransfersOfDeposit(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.ToAccountID, data.toAccount.ID)
	}
}

func TestListTransfersOfWithdraw(t *testing.T) {
	data := CreateRandomTransferParams{
		toAccount:   createRandomAccount(t),
		fromAccount: createRandomAccount(t),
	}
	for i := 0; i < 10; i++ {
		createRandomTransfer(t, data)
	}

	arg := ListTransfersOfWithdrawParams{FromAccountID: data.fromAccount.ID, Limit: 5, Offset: 5}

	transfers, err := testQueries.ListTransfersOfWithdraw(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.FromAccountID, data.fromAccount.ID)
	}
}
