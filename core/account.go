package core

import (
	"strings"

	"github.com/rs/zerolog/log"
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
	Id         string
	Username   string
	Characters []string
	Password   string
	Banned     bool
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

func (a *Account) GetID() string {
	return a.Id
}

// GetUsername implements Account.
func (a *Account) GetUsername() string {
	return a.Username
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

// Save implements Account.
func (a *Account) Save(callback func()) {
	panic("unimplemented")
}

// Serialize implements Account.
func (a *Account) Serialize() string {
	return ""
}

// SetPassword implements Account.
func (a *Account) SetPassword(password string) {
	a.Password = password
}

// UndeleteCharacter implements Account.
func (a *Account) UndeleteCharacter(name string) {
	panic("unimplemented")
}
