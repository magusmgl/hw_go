package bins

import (
	"errors"
	"strings"
	"time"
)

type Db interface {
	Read() ([]Bin, error)
	Write([]Bin) error
}

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
	Db   Db
}

func NewBin(id string, private bool, name string) (*Bin, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("EMPTY_ID")
	}
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("EMPTY_NAME")
	}

	newBin := &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}

	return newBin, nil
}

func NewBinList(db Db) *BinList {
	data, err := db.Read()
	if err != nil {
		return &BinList{
			Bins: []Bin{},
			Db:   db,
		}
	}
	return &BinList{
		Bins: data,
		Db:   db,
	}
}
