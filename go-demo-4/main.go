package main

import (
	"demo/password/account"
	"demo/password/encrypter"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"strings"

	"github.com/joho/godotenv"

	"github.com/fatih/color"
)

var memuVariants = []string{
	"1 - Создать аккаунт",
	"2 - Найти аккаунт по url",
	"3 - Найти аккаунт по логину",
	"4 - Удалить аккауунт",
	"5 - Выход",
	"Введите команду",
}
var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func main() {
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Не удалось найти env")
	}

	// vault := account.NewVault(cloud.NewCloudDB("https://ya.ru"))
	vault := account.NewVault(files.NewJsonDb("data.vault"), encrypter.NewEncrypter())
	fmt.Println("Приложение для хранения паролей")
Menu:
	for {
		variant := promptData(memuVariants...)
		menuFunction := menu[variant]
		if menuFunction == nil {
			break Menu
		}
		menuFunction(vault)
		// switch variant {
		// case "1":
		// 	createAccount(vault)
		// case "2":
		// 	findAccount(vault)
		// case "3":
		// 	deleteAccount(vault)
		// default:
		// 	break Menu
		// }
	}

}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData("Введите url для поиска")
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})

	outputResult(&accounts)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData("Введитие логин для поиска")
	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})

	outputResult(&accounts)
}

func outputResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		color.Red("Аккаунту с таким логином не найдены")
	}

	fmt.Println("Найденные аккаунты: ")
	for _, accounts := range *accounts {
		accounts.OutputAccount()
	}
}

// func checkLogin(acc account.Account, str string) bool {
// 	return strings.Contains(acc.Login, str)
// }

// func chekckUrl(acc account.Account, str string) bool {
// 	return strings.Contains(acc.Url, str)
// }

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите url для поиска: ")
	isDeleted := vault.DeleteAccounts(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		color.Red("Не найдено")
	}

}

func createAccount(vault *account.VaultWithDb) {
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

func promptData(prompt ...string) string {
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
