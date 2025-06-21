package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"os-test/bins"
)

func SaveStorage(fileName string, b *bins.Bin) {
	binsList := ReadFile(fileName)

	file, err := os.Create("test.json")

	if err != nil {
		fmt.Println("не удалость открыть файл")
	} 

	defer file.Close()
	binsList.AddBin(*b)
	fmt.Printf("binsList: %v\n", binsList.Bins)

	_, err = file.WriteString(string(binsList.ToBytes()))

	if err != nil {
		fmt.Println(err)
	}

}

func ReadFile(fileName string) *bins.BinList {
	binsList := bins.BinList{Bins: []bins.Bin{}}
	strBytes, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("не удалост прочеть файл")
	}

	err = json.Unmarshal(strBytes, &binsList)

	if err != nil {
		fmt.Println("не удалось заэнкодить")
	}

	return &binsList
}
