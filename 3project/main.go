package main

import (
	"errors"
	"strings"
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

func newBin(id string, private bool, name string) (*Bin, error) {
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
	bins []*Bin
}

func newBinList() *BinList {
	return &BinList{
		bins: make([]*Bin, 0),
	}
}

func main() {

}
