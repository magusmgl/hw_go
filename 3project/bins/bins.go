package bins

import (
	"errors"
	"strings"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
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

type BinList struct {
	Bins []*Bin `json:"bins"`
}

func NewBinList() *BinList {
	return &BinList{
		Bins: make([]*Bin, 0),
	}
}
