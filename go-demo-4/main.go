package main

import (
	"demo/password/account"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	vault := account.NewVault()
	fmt.Println("Приложение для хранения паролей")
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}

}

func getMenu() int {
	var variant int
	fmt.Println("Введите команду: ")
	fmt.Println("1 - Создать аккаунт")
	fmt.Println("2 - Найти аккаунт")
	fmt.Println("3 - Удалить аккауунт")
	fmt.Println("4 - Выход")
	fmt.Scan(&variant)
	return variant
}

func findAccount(vault *account.Vault) {
	url := promptData("Введите url для поиска")
	accounts := vault.FindAccountsByUrl(url)

	if len(accounts) == 0 {
		color.Red("Аккаунту с таким URL не найдены")
	}

	fmt.Println("Найденные аккаунты: ")
	for _, accounts := range accounts {
		accounts.OutputAccount()
	}
}

func deleteAccount(vault *account.Vault) {
	url := promptData("Введите url для поиска: ")
	isDeleted := vault.DeleteAccounts(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		color.Red("Не найдено")
	}

}

func createAccount(vault *account.Vault) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формта url или логин")
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
