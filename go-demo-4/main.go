package main

import (
	"fmt"
	"math/rand/v2"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc *account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(int(len(letterRunes)))]
	}
	acc.password = string(res)
}

func newAccount(login string, password string, url string) *account {
	return &account{
		login:    login,
		password: password,
		url:      url,
	}
}

var letterRunes = []rune("abcdefghijklmnoprstuwvyzABCDEFJHOQPRSTYUXYWZ1234567890-*!")

func main() {
	// a := 5
	// fmt.Println(&a)
	// double(&a)
	// fmt.Println(a)

	// b := [4]int{1, 2, 3, 4}
	// reverse(&b)
	// fmt.Println(b)
	str := []rune("Привет)")
	for _, ch := range string(str) {
		fmt.Println(ch, string(ch))
	}

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount := newAccount(login, password, url)
	myAccount.generatePassword(12)
	myAccount.outputPassword()

}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}

func double(num *int) {
	*num = *num * 2
}

func reverse(arr *[4]int) {
	for index, value := range *arr {
		(*arr)[len(arr)-1-index] = value
	}
}
