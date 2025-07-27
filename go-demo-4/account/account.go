package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"strings"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnoprstuwvyzABCDEFJHOQPRSTYUXYWZ1234567890-*!")

type Account struct {
	Login      string    `json:"login"`
	Password   string    `json:"password"`
	Url        string    `json:"url"`
	CreateTime time.Time `json:"createTime"`
<<<<<<< Updated upstream
	UpdateTine time.Time `json:"updateTime"`
}

func (acc *Account) OutputAccount() {
	color.Cyan(acc.Login)
	color.Cyan(acc.Password)
	color.Cyan(acc.Url)
=======
	UpdateTime time.Time `json:"updateTime"`
}

func (acc *Account) OutputAccount() {
	color.Cyan("  Login: %s", acc.Login)
	color.Cyan("  Password: %s", acc.Password)
	color.Cyan("  URL: %s", acc.Url)
>>>>>>> Stashed changes
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(int(len(letterRunes)))]
	}
	acc.Password = string(res)
}

func NewAccount(login string, password string, urlString string) (*Account, error) {
	if strings.TrimSpace(login) == "" {
		return nil, errors.New("EMPTY_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &Account{
		CreateTime: time.Now(),
<<<<<<< Updated upstream
		UpdateTine: time.Now(),
=======
		UpdateTime: time.Now(),
>>>>>>> Stashed changes
		Login:      login,
		Password:   password,
		Url:        urlString,
	}
	if strings.TrimSpace(password) == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
