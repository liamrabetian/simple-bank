package db

import (
	"context"
	"math/rand"
	"testing"

	"github.com/mohammadrabetian/simple-bank/utils"
	"github.com/stretchr/testify/require"
)

func RandomOwnerName() string {
	return utils.RandomString(8)
}

func RandomMoneyAmount() int64 {
	return utils.RandomInt(1, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func CreateRandomAccount(t *testing.T, owner ...string) Account {
	ownerValue := RandomOwnerName()

	if len(owner) > 0 {
		ownerValue = owner[0]
	}
	arg := CreateAccountParams{Owner: ownerValue, Balance: RandomMoneyAmount(), Currency: RandomCurrency()}
	a, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, a)
	require.Equal(t, a.Balance, arg.Balance)
	require.Equal(t, a.Currency, arg.Currency)
	require.NotZero(t, a.CreatedAt)
	return a
}
