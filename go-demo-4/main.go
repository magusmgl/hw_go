package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, error := account.NewAccountWithTimestamp(login, password, url)
	if error != nil {
		fmt.Println("Неверный формта url или логин")
		return
	}
	myAccount.OutputPassword()
	files.WriteFile()
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
