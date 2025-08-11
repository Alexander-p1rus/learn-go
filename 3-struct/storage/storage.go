package storage

import (
	"encoding/json"
	"errors"
	"os"

	"os-test/bins"
)

func SaveStorage(fileName string, b *bins.Bin) error {

	file, err := os.Open(fileName)

	if err != nil {
		file, err = os.Create(fileName)
		if err != nil {
			return errors.New("не удалось создать файл")
		}

	}

	binsList, err := ReadFile(fileName)
	if err != nil {
		return err
	}

	binsList.AddBin(b)

	err = os.WriteFile(fileName, binsList.ToBytes(), 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}

func ReadFile(fileName string) (*bins.BinList, error) {
	binsList := bins.BinList{Bins: []bins.Bin{}}
	strBytes, err := os.ReadFile(fileName)

	if err != nil {
		return nil, errors.New("не удалось прочесть файл")
	}

	if len(strBytes) == 0 {
		return &binsList, nil
	}

	err = json.Unmarshal(strBytes, &binsList)

	if err != nil {
		return nil, errors.New("не удалость считать json файл")
	}

	return &binsList, nil
}
