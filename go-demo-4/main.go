package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"strings"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimestamp struct {
	createTime time.Time
	updateTine time.Time
	account
}

func (acc *accountWithTimestamp) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url, acc.createTime, acc.updateTine)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(int(len(letterRunes)))]
	}
	acc.password = string(res)
}

func newAccountWithTimestamp(login string, password string, urlString string) (*accountWithTimestamp, error) {
	if strings.TrimSpace(login) == "" {
		return nil, errors.New("EMPTY_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &accountWithTimestamp{
		createTime: time.Now(),
		updateTine: time.Now(),
		account: account{
			login:    login,
			password: password,
			url:      urlString,
		},
	}
	if strings.TrimSpace(password) == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil

}

// func newAccount(login string, password string, urlString string) (*account, error) {
// 	if strings.TrimSpace(login) == "" {
// 		return nil, errors.New("EMPTY_LOGIN")
// 	}

// 	_, err := url.ParseRequestURI(urlString)
// 	if err != nil {
// 		return nil, errors.New("INVALID_URL")
// 	}

// 	newAcc := &account{
// 		login:    login,
// 		password: password,
// 		url:      urlString,
// 	}

// 	if strings.TrimSpace(password) == "" {
// 		newAcc.generatePassword(12)
// 	}
// 	return newAcc, nil
// }

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

	myAccount, error := newAccountWithTimestamp(login, password, url)
	if error != nil {
		fmt.Println("Неверный формта url или логин")
		return
	}
	// myAccount.generatePassword(12)
	myAccount.outputPassword()

}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
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
