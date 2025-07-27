package file

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func ReadFile(fileName string) ([]byte, error) {
	if !IsJsonFile(fileName) {
		return nil, errors.New("NOT_JSON_FILE")
	}
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteFile(contet []byte, fileName string) {
	if !IsJsonFile(fileName) {
		color.Red("Указанный файл не является JSON файлом")
		return
	}
	file, err := os.Create(fileName)
	if err != nil {
		color.Red("Не удалось создать файл")
		return
	}
	defer file.Close()
	_, err = file.Write(contet)
	if err != nil {
		color.Red("Не удалось записать данные в файл")
		return
	}
	color.Green("Данные записаны в файл")
}

func IsJsonFile(filename string) bool {
	return filepath.Ext(filename) != ".json"
}
