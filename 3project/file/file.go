package file

import (
	"os"

	"github.com/fatih/color"
)

func ReadFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		color.Red("Не удалось прочитать файл")
		return nil, err
	}
	return data, nil
}

func WriteFile(contet []byte, fileName string) {
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
