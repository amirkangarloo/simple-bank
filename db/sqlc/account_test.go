package db

import (
	"context"
	"database/sql"
	"simple-bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, account.Owner, arg.Owner)
	require.Equal(t, account.Balance, arg.Balance)
	require.Equal(t, account.Currency, arg.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	fetchAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, fetchAccount)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)
	fetchAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchAccount)

	require.Equal(t, account.ID, fetchAccount.ID)
	require.Equal(t, account.Owner, fetchAccount.Owner)
	require.Equal(t, account.Balance, fetchAccount.Balance)
	require.Equal(t, account.Currency, fetchAccount.Currency)
	require.WithinDuration(t, account.CreatedAt, fetchAccount.CreatedAt, time.Second)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{Limit: 5, Offset: 5}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}
	fetchAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, fetchAccount)

	require.Equal(t, account.ID, fetchAccount.ID)
	require.Equal(t, arg.Balance, fetchAccount.Balance)
	require.Equal(t, account.Owner, fetchAccount.Owner)
	require.Equal(t, account.Currency, fetchAccount.Currency)
	require.WithinDuration(t, account.CreatedAt, fetchAccount.CreatedAt, time.Second)
}
