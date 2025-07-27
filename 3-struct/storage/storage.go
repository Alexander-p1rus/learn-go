package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"os-test/bins"
)

func SaveStorage(fileName string, b *bins.Bin) error{
	binsList, err := ReadFile(fileName)

	if err != nil {
		return err
	}
	file, err := os.Create("test.json")

	if err != nil {
		return errors.New("не удалось создать файл")
	} 

	defer file.Close()
	binsList.AddBin(*b)
	fmt.Printf("binsList: %v\n", binsList.Bins)

	_, err = file.WriteString(string(binsList.ToBytes()))

	if err != nil {
		return err
	}

	return nil
}

func ReadFile(fileName string) (*bins.BinList , error) {
	binsList := bins.BinList{Bins: []bins.Bin{}}
	strBytes, err := os.ReadFile(fileName)

	if err != nil {
		return nil,errors.New("не удалось прочесть файл")
	}

	err = json.Unmarshal(strBytes, &binsList)

	if err != nil {
		return nil , errors.New("не удалость считать json файл")
	}

	return &binsList, nil}
