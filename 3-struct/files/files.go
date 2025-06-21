package files

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadFile(fileName string) ([]byte, error) {
	isJson := isJson(fileName)

	if !isJson {
		return nil, errors.New("не json формат файла")
	}

	data, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("не удалось прочесть файл")
	}

	return data	, nil

}

func isJson(str string) bool {
	return strings.Contains(str, ".json")
}
