package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	TypesOperation := [3]string{"AVG", "SUM", "MED"}

	opertaionType := inputOperationType(TypesOperation[:])
	valueList := getOperationValue()

	result := calculateByOperation(opertaionType, valueList)
	fmt.Printf("result: %v\n", result)
}

func inputOperationType(types []string) string {
	fmt.Println("введите тип операции:AVG/SUM/MED")
	var inputType string

Input:
	for {
		fmt.Scanln(&inputType)

		for _, value := range types {
			if value == strings.ToUpper(inputType) {
				break Input
			}
		}

		fmt.Println("Введите корректное название операции")
	}

	return strings.ToUpper(inputType)
}

func getOperationValue() []string {
	var str string
	fmt.Println("Введите через запятую целые числа")
	fmt.Scan(&str)

	return strings.Split(str, ",")

}

func calculateByOperation(typeOp string, valueL []string) float64 {

	switch typeOp {
	case "AVG":
		return calcAvarage(valueL)
	case "SUM":
		return calcSum(valueL)
	case "MED":
		return calcMediana(valueL)
	}

	return 0
}

func calcAvarage(List []string) float64 {
	var total int
	for _, value := range List {
		parseValue, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("что-то пошло не так")
		}

		total += int(parseValue)
	}

	return float64(total) / float64(len(List))

}

func calcSum(List []string) float64 {
	sum := 0

	for _, listValue := range List {
		number, err := strconv.Atoi(listValue)
		if err != nil {
			fmt.Println("Парсинг в тип данных Int не удался")
		}
		sum += number
	}

	return float64(sum)

}

func calcMediana(List []string) float64 {
	newArr := make([]int, 0, cap(List))

	for _, value := range List {
		valueToInt, _ := strconv.Atoi(value)
		newArr = append(newArr, valueToInt)
	}

	sort.Ints(newArr)

	if len(newArr)%2 == 0 {
		num1:= float64(newArr[(len(newArr)-1)/2])
		num2:= float64(newArr[len(newArr)/2])
		return (num1+num2)/2

	} else {
		return float64(newArr[(len(newArr)-1)/2])
	}

}
