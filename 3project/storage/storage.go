package storage

import (
	"3/cli/bins"
	"3/cli/file"
	"encoding/json"
)

func SaveBinsToFile(fileName string, binList *bins.BinList) error {
	content, err := json.Marshal(binList)
	if err != nil {
		return err
	}

	err = file.WriteToJsonFile(content, fileName)
	if err != nil {
		return err
	}
	return nil
}

func ReadBinsFromFile(fileName string, binList *bins.BinList) error {
	data, err := file.ReadJsonFile(fileName)
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
