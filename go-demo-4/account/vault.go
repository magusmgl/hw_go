package account

import (
	"demo/password/output"
	"encoding/json"
	"strings"
	"time"
)

type FromByteDb interface {
	Read() ([]byte, error)
}

type ToByteDb interface {
	Write([]byte)
}

type MyFunc func(a string) bool

type Db interface {
	FromByteDb
	ToByteDb
}

type Encrypter interface {
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
}

type Vault struct {
	Accounts []Account `json:"accounts"`
	UpdateAt time.Time `json:"updateAt"`
}

type VaultWithDb struct {
	Vault
	db  Db
	enc Encrypter
}

func NewVault(db Db, enc Encrypter) *VaultWithDb {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts: []Account{},
				UpdateAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	data := enc.Decrypt(file)
	var vault Vault
	err = json.Unmarshal(data, &vault)
	if err != nil {
		output.PrintError("Не удалось разобрать файл data.vault")
		return &VaultWithDb{
			Vault: Vault{
				Accounts: []Account{},
				UpdateAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	return &VaultWithDb{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *VaultWithDb) FindAccounts(str string, checker func(Account, string) bool) []Account {
	var foundAccounts []Account
	for _, account := range vault.Accounts {
		isMatched := checker(account, str)
		if isMatched {
			foundAccounts = append(foundAccounts, account)
		}
	}
	return foundAccounts
}

func (vault *VaultWithDb) DeleteAccounts(url string) bool {
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

func (vault *VaultWithDb) save() {
	vault.UpdateAt = time.Now()
	data, err := vault.Vault.ToBytes()
	encData := vault.enc.Encrypt(data)
	if err != nil {
		output.PrintError("Не удалось преобразовать")
	}
	vault.db.Write(encData)
}
