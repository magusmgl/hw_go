package storage

import (
	"3/cli/bins"
	"3/cli/file"
	"encoding/json"
	"errors"
	"os"
)

type Storagedb struct {
	FileName string
}

func NewStorageDb(filename string) *Storagedb {
	return &Storagedb{
		FileName: filename,
	}
}

func (db *Storagedb) Write(binList []bins.Bin) error {
	content, err := json.Marshal(binList)
	if err != nil {
		return err
	}

	file, err := os.Create(db.FileName)

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

func (db *Storagedb) Read() ([]bins.Bin, error) {
	data, err := file.ReadJsonFile(db.FileName)
	if err != nil {
		return nil, errors.New("FAILED_TO_OPEN_FILE")
	}
	if len(data) == 0 {
		return nil, errors.New("NO_DATA_TO_READ")
	}

	var bins []bins.Bin
	err = json.Unmarshal(data, &bins)

	if err != nil {
		return nil, errors.New("ERROR")
	}
	return bins, nil
}
