package files

import (
	"fmt"
	"os"
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
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err1 := file.Write(content)
	if err1 != nil {

		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}
