package main

import (
	"fmt"
	"time"

	"os-test/bins"
	"os-test/storage"
)

func main() {

	bin, err := bins.СreateBin("222", true, time.Now(), "hello")

	if err != nil {
		fmt.Println(err.Error())
	}

	storage, err := storage.CreateStorage("storage.json")
	if err != nil {
		fmt.Println(err.Error())

	}

	bl, err := storage.Read(bins.BinList{})
	if err != nil {
		return
	}

	bl.AddBin(bin)

	err = storage.Save(bl.ToBytes())
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
	}

}
