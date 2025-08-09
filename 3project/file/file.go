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

func isJsonFile(filename string) bool {
	return filepath.Ext(filename) == ".json"
}
