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
		panic(err)
	}

	storage.SaveStorage("test.json", bin)

	bytes, _ := files.ReadFile("test.json")
	fmt.Printf("bytes: %v\n", string(bytes))
}
