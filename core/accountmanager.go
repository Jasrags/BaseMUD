package core

import (
	"errors"

	"github.com/rs/zerolog/log"
)

// type AccountManager interface {
// 	AddAccount(acc Account) error
// 	GetAccount(username string) (Account, error)
// 	LoadAccount(username string, force bool) (Account, error)
// 	LoadAccounts() error
// }

func NewAccountManager() *AccountManager {
	return &AccountManager{
		Accounts: make(map[string]*Account),
	}
}

type AccountManager struct {
	Accounts map[string]*Account
}

// AddAccount implements AccountManager.
func (a *AccountManager) AddAccount(acc *Account) error {
	// Check if the account is already registered, error if so
	if _, ok := a.Accounts[acc.GetUsername()]; ok {
		log.Error().Msg("Account already exists")
		return errors.New("account already exists")
	}

	// Add the account to the map
	a.Accounts[acc.GetUsername()] = acc

	return nil
}

// GetAccount implements AccountManager.
func (a *AccountManager) GetAccount(username string) (*Account, error) {
	if acc, ok := a.Accounts[username]; ok {
		return acc, nil
	}

	log.Warn().Msg("Account not found")
	return nil, errors.New("account not found")
}

// LoadAccount implements AccountManager.
func (a *AccountManager) LoadAccount(username string, force bool) (*Account, error) {
	panic("unimplemented")
}

func (a *AccountManager) LoadAccounts() error {
	log.Info().Msg("Loading accounts")
	a.Accounts = map[string]*Account{
		"test": {
			Id:         "1",
			Username:   "test",
			Characters: []string{"test"},
			Password:   "test",
			Banned:     false,
		},
	}

	return nil
}
