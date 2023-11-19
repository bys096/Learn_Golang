package main

import (
	"fmt"

	"github.com/bys096/banking/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(500)
	err := account.Withdraw(1500)
	if err != nil {
		// log.Fatalln(err) // 에러를 출력하고, 프로그램을 종료시킴
		// fmt.Println(err)
	}
	// fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}
