package core

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// type Account interface {
// 	AddCharacter(username string)
// 	Ban()
// 	CheckPassword(pass string) bool
// 	DeleteAccount()
// 	DeleteCharacter(name string)
// 	GetID() string
// 	GetUsername() string
// 	HasCharacter(name string) bool
// 	Save(callback func())
// 	Serialize() string
// 	SetPassword(password string)
// 	UndeleteCharacter(name string)
// }

type Account struct {
	Id         string   `json:"id"`
	Username   string   `json:"username"`
	Characters []string `json:"characters"`
	Password   string   `json:"password"`
	Banned     bool     `json:"banned"`
}

func NewAccount() *Account {
	return &Account{}
}

// AddCharacter implements Account.
func (a *Account) AddCharacter(username string) {
	a.Characters = append(a.Characters, username)
}

// Ban implements Account.
func (a *Account) Ban() {
	a.Banned = true
}

// CheckPassword implements Account.
func (a *Account) CheckPassword(pass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(pass)); err != nil {
		log.Error().Err(err).Msg("Failed to compare password")
		return false
	}

	return true
}

// DeleteAccount implements Account.
func (a *Account) DeleteAccount() {
	panic("unimplemented")
}

// DeleteCharacter implements Account.
func (a *Account) DeleteCharacter(name string) {
	for i, v := range a.Characters {
		if strings.EqualFold(v, name) {
			a.Characters = append(a.Characters[:i], a.Characters[i+1:]...)
			return
		}
	}
}

// HasCharacter implements Account.
func (a *Account) HasCharacter(name string) bool {
	for _, v := range a.Characters {
		if strings.EqualFold(v, name) {
			return true
		}
	}

	return false
}

func (a *Account) Save() {
	var sb strings.Builder

	e := json.NewEncoder(&sb)
	e.SetEscapeHTML(false)
	e.SetIndent("", "  ")

	if err := e.Encode(a); err != nil {
		log.Error().
			Err(err).
			Msg("Failed to encode account to JSON")
		return
	}

	filePath := fmt.Sprintf("%s/%s.json",
		viper.GetString("data.accounts_path"), strings.ToLower(a.Username))

	if err := os.WriteFile(filePath, []byte(sb.String()), 0644); err != nil {
		log.Error().
			Err(err).
			Msg("Failed to write account to file")
		return
	}
}

// SetPassword implements Account.
func (a *Account) SetPassword(password string) {
	a.Password = password
}

// UndeleteCharacter implements Account.
func (a *Account) UndeleteCharacter(name string) {
	panic("unimplemented")
}
