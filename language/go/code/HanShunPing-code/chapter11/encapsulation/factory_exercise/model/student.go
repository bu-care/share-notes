package model

import (
	"fmt"
)

type account struct {
	accountId string
	password  string
	balance   float64
}

func NewAccount(accountId string, password string, balance float64) *account {
	idLen := len(accountId)
	if idLen < 6 || idLen > 10 {
		fmt.Println("length of accountId is error.")
		return nil
	}
	if len(password) != 6 {
		fmt.Println("length of password is incorrect.")
		return nil
	}
	if balance < 20 {
		fmt.Println("balance is incorrect.")
		return nil
	}
	return &account{
		accountId: accountId,
		password:  password,
		balance:   balance,
	}
}

func (a *account) SetFeild(accountId string, password string, balance float64) {
	a.accountId = accountId
	a.password = password
	a.balance = balance
}

func (a *account) GetAccountInfo() (string, string, float64) {
	return a.accountId, a.password, a.balance
}
