package storage

import (
	"3/cli/bins"
	"encoding/json"
	"errors"
	"os"
)

type Storagedb struct {
	filename string
}

func NewStorageDb(filename string) *Storagedb {
	return &Storagedb{
		filename: filename,
	}
}

func (storageDb *Storagedb) Write(binList *bins.BinList) error {
	content, err := json.Marshal(binList)
	if err != nil {
		return err
	}

	file, err := os.Create(storageDb.filename)

	if err != nil {
		return errors.New("FAILED_TO_CREATE_FILE")
	}

	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		return errors.New("FALIED_TO_WRITE_DATA")
	}

	return nil
}

func (storageDb *Storagedb) ReadBinsFromFile() (error, *bins.BinList) {
	data, err := os.ReadFile(storageDb.filename)
	if err != nil {
		return errors.New("FAILED_TO_OPEN_FILE"), nil
	}
	if len(data) == 0 {
		return errors.New("NO_DATA_TO_READ"), nil
	}

	binList := bins.NewBinList()
	err = json.Unmarshal(data, binList)
	if err != nil {
		return errors.New("FAILED_CONVERT_DATA_TO_BIN"), nil
	}
	return nil, binList
}
