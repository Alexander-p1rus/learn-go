package storage

import (
	"encoding/json"
	"errors"
	"os"

	"os-test/bins"
)

func Save(fileName string, bl *bins.BinList) error {
	err := os.WriteFile(fileName, bl.ToBytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func Read(fileName string) (*bins.BinList, error) {
	binsList := bins.BinList{Bins: []bins.Bin{}}

	strBytes, err := os.ReadFile(fileName)

	if err != nil {
		file, err := os.Create(fileName)
		if err != nil {
			return nil, err
		}

		file.Close()

		return &binsList, nil
	}

	err = json.Unmarshal(strBytes, &binsList)
	if err != nil {
		return nil, errors.New("не удалость считать json файл")
	}

	return &binsList, nil
}
