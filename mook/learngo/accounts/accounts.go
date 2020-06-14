package accounts

import "fmt"

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	fmt.Println("account addr : ", &account)
	fmt.Println("account val : ", *(&account))
	return &account
}

// a : reciever -> struct 이름 첫글자의 소문자
// Deposit x amout on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}
