package files

import (
	"demo/password/output"
	"os"

	"github.com/fatih/color"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		output.PrintError(err)
		return nil, err
	}

	return data, nil
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		output.PrintError(err)
		return
	}
	defer file.Close()
	_, err1 := file.Write(content)
	if err1 != nil {

		output.PrintError(err)
		return
	}
	color.Green("Запись успешна")
}
