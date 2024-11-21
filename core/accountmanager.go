package core

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
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
func (am *AccountManager) AddAccount(acc *Account) error {
	// Check if the account is already registered, error if so
	if _, ok := am.Accounts[acc.Username]; ok {
		log.Error().Msg("Account already exists")
		return errors.New("account already exists")
	}

	// Add the account to the map
	am.Accounts[acc.Username] = acc

	return nil
}

// GetAccount implements AccountManager.
func (am *AccountManager) GetAccount(username string) (*Account, error) {
	if acc, ok := am.Accounts[strings.ToLower(username)]; ok {
		return acc, nil
	}

	log.Warn().
		Str("username", username).
		Msg("Account not found")
	return nil, errors.New("account not found")
}

// LoadAccount implements AccountManager.
func (am *AccountManager) LoadAccount(username string, force bool) (*Account, error) {
	panic("unimplemented")
}

func (am *AccountManager) LoadAccounts() error {
	log.Info().Msg("Loading accounts")

	files, err := os.ReadDir(viper.GetString("data.accounts_path"))
	if err != nil {
		log.Error().Err(err).Msg("Failed to read accounts directory")

		return err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join("_data/accounts", file.Name())

			data, err := os.ReadFile(filePath)
			if err != nil {
				log.Error().Err(err).Msgf("Failed to read file %s", filePath)
				continue
			}

			var acc Account
			if err := json.Unmarshal(data, &acc); err != nil {
				log.Error().Err(err).Msgf("Failed to unmarshal account data from file %s", filePath)
				continue
			}

			am.Accounts[strings.ToLower(acc.Username)] = &acc
			log.Debug().
				Str("username", acc.Username).
				Msg("Loaded account")
		}
	}

	log.Info().
		Int("count", len(am.Accounts)).
		Msg("Accounts loaded")

	return nil
}
