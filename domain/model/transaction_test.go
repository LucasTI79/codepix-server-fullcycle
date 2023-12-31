package model_test

import (
	"testing"

	"github.com/lucasti79/codepix-go/domain/model"
	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/require"
)

func TestNewTransaction(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := model.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Wesley"
	account, _ := model.NewAccount(bank, accountNumber, ownerName)

	accountNumberDestination := "abcdestination"
	ownerName = "Mariana"
	accountDestination, _ := model.NewAccount(bank, accountNumberDestination, ownerName)

	kind := "email"
	key := "j@j.com"
	pixKey, _ := model.NewPixKey(kind, accountDestination, key)

	require.NotEqual(t, account.ID, accountDestination.ID)

	amount := 3.10
	statusTransaction := "pending"
	transaction, err := model.NewTransaction(account, amount, pixKey, "My description", "")
	//
	require.Nil(t, err)
	require.NotNil(t, uuid.FromStringOrNil(transaction.ID))
	require.Equal(t, amount, transaction.Amount)
	require.Equal(t, statusTransaction, transaction.Status)
	require.Equal(t, "My description", transaction.Description)
	require.Empty(t, transaction.CancelDescription)

	pixKeySameAccount, err := model.NewPixKey(kind, account, key)

	_, err = model.NewTransaction(account, amount, pixKeySameAccount, "My description", "")
	require.NotNil(t, err)

	_, err = model.NewTransaction(account, 0, pixKey, "My description", "")
	require.NotNil(t, err)
}

func TestModel_ChangeStatusOfATransaction(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := model.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Wesley"
	account, _ := model.NewAccount(bank, accountNumber, ownerName)

	accountNumberDestination := "abcdestination"
	ownerName = "Mariana"
	accountDestination, _ := model.NewAccount(bank, accountNumberDestination, ownerName)

	kind := "email"
	key := "j@j.com"
	pixKey, _ := model.NewPixKey(kind, accountDestination, key)

	amount := 3.10
	transaction, _ := model.NewTransaction(account, amount, pixKey, "My description", "")

	transaction.Complete()
	require.Equal(t, model.TransactionCompleted, transaction.Status)

	transaction.Cancel("Error")
	require.Equal(t, model.TransactionError, transaction.Status)
	require.Equal(t, "Error", transaction.CancelDescription)

}
