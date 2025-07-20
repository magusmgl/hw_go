package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"strings"
	"time"
)

var letterRunes = []rune("abcdefghijklmnoprstuwvyzABCDEFJHOQPRSTYUXYWZ1234567890-*!")

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

func (acc *accountWithTimestamp) OutputPassword() {
	fmt.Println(acc.login, acc.password, acc.url, acc.createTime, acc.updateTine)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(int(len(letterRunes)))]
	}
	acc.password = string(res)
}

func NewAccountWithTimestamp(login string, password string, urlString string) (*accountWithTimestamp, error) {
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
