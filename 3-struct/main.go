package main

import (
	"fmt"
	"time"

	"os-test/bins"
	"os-test/files"
	"os-test/storage"
)

func main() {

	bin, err := bins.СreateBin("222", true, time.Now(), "hello")
	if err != nil {
		fmt.Print(err.Error())
	}

	binsList, err := storage.Read("test.json")
	fmt.Printf("binsListAFterRead: %v\n", binsList)
	if err != nil {
		fmt.Println(err.Error())
	}
	binsList.AddBin(bin)
	fmt.Printf("binsList: %v\n", binsList)

	storage.Save("test.json", binsList)

	file, err := files.ReadFile("test.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("file: %v\n", string(file))

}
