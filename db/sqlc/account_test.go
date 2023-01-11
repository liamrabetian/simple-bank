package db

import (
	"context"
	"testing"

	"github.com/mohammadrabetian/simple-bank/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	createdAccount := CreateRandomAccount(t)
	gotAccount, err := testQueries.GetAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gotAccount)
	require.Equal(t, createdAccount.ID, gotAccount.ID)
}

func TestListAccounts(t *testing.T) {
	owner := utils.RandomString(8)
	for i := 0; i < 5; i++ {
		CreateRandomAccount(t, owner)
	}
	arg := ListAccountsParams{Owner: owner, Limit: 10, Offset: 1}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 4)
	for _, a := range accounts {
		require.Equal(t, a.Owner, owner)
	}
}

func TestUpdateAccount(t *testing.T) {
	account := CreateRandomAccount(t)
	newBalance := utils.RandomInt(account.Balance+1, account.Balance*2)
	arg := UpdateAccountParams{ID: account.ID, Balance: newBalance}
	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEqual(t, account.Balance, updatedAccount.Balance)
	require.Equal(t, updatedAccount.Balance, newBalance)
	require.Equal(t, account.Currency, updatedAccount.Currency)
}

func TestDeleteAccount(t *testing.T) {
	account := CreateRandomAccount(t)
	testQueries.DeleteAccount(context.Background(), account.ID)
	deletedAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.Empty(t, deletedAccount)
}
