package main

import (
	"3/cli/bins"
	"3/cli/storage"
	"fmt"
)

func main() {
	db := storage.NewStorageDb("data.json")
	binList := bins.NewBinList(db)
	bin1, _ := bins.NewBin("21312", false, "dsada")
	bins := make([]bins.Bin, 1)
	bins = append(bins, *bin1)
	err := binList.Db.Write(&bins)
	if err != nil {
		fmt.Println(err.Error())
	}
	bins2, err := binList.Db.Read()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(bins2)
}
