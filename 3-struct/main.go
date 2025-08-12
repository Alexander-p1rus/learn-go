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
		fmt.Println(err.Error())
	}

	storage.SaveStorage("test.json", bin)

	file, err := files.ReadFile("test.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("file: %v\n", string(file))

}
