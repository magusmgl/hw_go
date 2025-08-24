package main

import (
	"3/cli/api"
	"3/cli/bins"
	"3/cli/config"
	"3/cli/storage"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Не удалост найти env")
	}

	api.NewAPI(*config.NewConfig())

	binList := bins.NewBinList(storage.NewStorageDb("data.json"))

	bin1, _ := bins.NewBin("21312", false, "dsada")
	bins := make([]bins.Bin, 1)
	bins = append(bins, *bin1)

	err = binList.Db.Write(bins)
	if err != nil {
		fmt.Println(err.Error())
	}
	bins2, err := binList.Db.Read()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(bins2)
}
