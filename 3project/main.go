package main

import (
	"errors"
	"strings"
	"time"
)

type Bin struct {
	id      string
	private bool
	created time.Time
	name    string
}

func newBin(id string, private bool, name string) (*Bin, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("EMPTY_ID")
	}
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("EMPTY_NAME")
	}

	newBin := &Bin{
		id:      id,
		private: private,
		created: time.Now(),
		name:    name,
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
