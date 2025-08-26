package storage

import (
	"encoding/json"
	"errors"
	"os"

	"os-test/bins"
)

type Storage struct {
	filename string
}

func (s *Storage) Save(data []byte) error {
	err := os.WriteFile(s.filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Read(bl bins.BinList) (bins.BinList, error) {
	strBytes, err := os.ReadFile(s.filename)
	if err != nil {
		return bl, errors.New("не удалось открыть файл")
	}

	if len(strBytes) == 0 {
		return bl, nil
	}

	err = json.Unmarshal(strBytes, &bl)
	if err != nil {
		return bl, errors.New("не удалость считать json файл")
	}

	return bl, nil
}

func CreateStorage(filename string) (*Storage, error) {
	storage := &Storage{
		filename: filename,
	}

	_, err := os.ReadFile(storage.filename)
	if err != nil {
		file, err := os.Create(storage.filename)
		if err != nil {
			return nil, err
		}

		defer file.Close()

	}

	return storage, nil
}
