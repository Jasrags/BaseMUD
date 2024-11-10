package core

type Account struct {
	// Username   string
	Characters []string
	// Password   string
	// Banned     bool
}

func NewAccount() *Account {
	return &Account{}
}

func (a *Account) AddCharacter(username string) {
}

func (a *Account) Ban() {

}

func (a *Account) CheckPassword(password string) bool {
	return false
}

func (a *Account) DeleteAccount() {

}

func (a *Account) DeleteCharacter(name string) {

}

func (a *Account) GetUsername() string {
	return ""
}

func (a *Account) HasCharacter(name string) bool {
	return false
}

func (a *Account) Save(callback func()) {

}

func (a *Account) SetPassword(password string) {

}

func (a *Account) UndeleCharacter(name string) {

}
