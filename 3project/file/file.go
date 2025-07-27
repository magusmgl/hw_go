package file

import (
	"errors"
	"os"
	"path/filepath"
)

func ReadJsonFile(fileName string) ([]byte, error) {
	if !isJsonFile(fileName) {
		return nil, errors.New("NOT_JSON_FILE")
	}
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteToJsonFile(contet []byte, fileName string) error {
	if !isJsonFile(fileName) {
		return errors.New("NOT_JSON_FILE")
	}

	file, err := os.Create(fileName)
	if err != nil {
		return errors.New("FAILED_TO_CREATE_FILE")
	}
	defer file.Close()
	_, err = file.Write(contet)
	if err != nil {
		return errors.New("FALIED_TO_WRITE_DATA")
	}
	return nil
}

func isJsonFile(filename string) bool {
	return filepath.Ext(filename) == ".json"
}
