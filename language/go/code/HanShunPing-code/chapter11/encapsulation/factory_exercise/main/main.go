package main

import (
	"fmt"
	"go_code/chapter11/encapsulation/factory_exercise/model"
)

func main() {
	account := model.NewAccount("12345678", "666666", 40)
	fmt.Println("Account init info:", account)

	account.SetFeild("32165488", "123456", 80)
	fmt.Println("Account set info:", account)

	id, password, balance := account.GetAccountInfo()
	fmt.Println("id, password, balance of account info:", id, password, balance)

}
