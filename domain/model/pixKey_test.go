package model_test

import (
	"testing"

	"github.com/lucasti79/codepix-go/domain/model"
	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/require"
)

func TestModel_NewPixKey(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, err := model.NewBank(code, name)
	require.Nil(t, err)

	accountNumber := "abcnumber"
	ownerName := "Wesley"
	account, err := model.NewAccount(bank, accountNumber, ownerName)
	require.Nil(t, err)

	kind := "email"
	key := "j@j.com"
	accountId := uuid.NewV4().String()
	pixKey, err := model.NewPixKey(kind, account, key, accountId)
	require.Nil(t, err)

	require.NotEmpty(t, uuid.FromStringOrNil(pixKey.ID))
	require.Equal(t, pixKey.Kind, kind)
	require.Equal(t, pixKey.Status, "active")

	kind = "cpf"
	_, err = model.NewPixKey(kind, account, key, accountId)
	require.Nil(t, err)

	_, err = model.NewPixKey("nome", account, key, accountId)
	require.NotNil(t, err)
}
