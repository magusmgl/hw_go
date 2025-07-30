package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	// vault := account.NewVault(cloud.NewCloudDB("https://ya.ru"))
	vault := account.NewVault(files.NewJsonDb("data.json"))
	fmt.Println("Приложение для хранения паролей")
Menu:
	for {
		variant := promptData([]string{
			"1 - Создать аккаунт",
			"2 - Найти аккаунт",
			"3 - Удалить аккауунт",
			"4 - Выход",
			"Введите команду",
		})
		switch variant {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}

}

func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите url для поиска"})
	accounts := vault.FindAccountsByUrl(url)

	if len(accounts) == 0 {
		color.Red("Аккаунту с таким URL не найдены")
	}

	fmt.Println("Найденные аккаунты: ")
	for _, accounts := range accounts {
		accounts.OutputAccount()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите url для поиска: "})
	isDeleted := vault.DeleteAccounts(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		color.Red("Не найдено")
	}

}

func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	url := promptData([]string{"Введите URL"})

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формта url или логин")
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData[T any](prompt []T) string {
	for i, val := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", prompt[i])
		} else {
			fmt.Println(val)
		}
	}
	var res string
	fmt.Scan(&res)
	return res
}
