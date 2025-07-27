package bins

import (
	"3/cli/file"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/fatih/color"
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

func (binList *BinList) SaveToBinsFile(fileName string) {
	content, err := json.Marshal(binList)
	if err != nil {
		color.Red("Не удалось преобразовать данные в json")
		return
	}
	file.WriteFile(content, fileName)
}

func (binList *BinList) ReadBinsFromFile(fileName string) error {
	data, err := file.ReadFile(fileName)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return nil
	}

	err = json.Unmarshal(data, binList)
	if err != nil {
		return err
	}
	return nil
}
