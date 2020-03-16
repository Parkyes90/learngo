package main

import (
	"fmt"
	"github.com/parkyes90/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("parkyes90")
	account.Deposit(500)
	err := account.Withdraw(6000)
	if err != nil {
		fmt.Println(err)
	}
	account.ChangeOwner("parkyes")

	fmt.Println(account.Balance(), account.Owner())
}
