package account

import (
	"demo/password/files"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts []Account `json:"accounts"`
	UpdateAt time.Time `json:"updateAt"`
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Accounts: []Account{},
			UpdateAt: time.Now(),
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
		return &Vault{
			Accounts: []Account{},
			UpdateAt: time.Now(),
		}
	}
	return &vault
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) FindAccountsByUrl(url string) []Account {
	var foundAccounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if isMatched {
			foundAccounts = append(foundAccounts, account)
		}
	}
	return foundAccounts
}

func (vault *Vault) DeleteAccounts(url string) bool {
	var newAcc []Account
	isDeleted := false
	for _, acc := range vault.Accounts {
		isMatched := strings.Contains(acc.Url, url)
		if !isMatched {
			newAcc = append(newAcc, acc)
			continue
		}
		isDeleted = true
	}
	vault.Accounts = newAcc
	vault.save()
	return isDeleted
}

func (vault *Vault) save() {
	vault.UpdateAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	files.WriteFile(data, "data.json")
}
