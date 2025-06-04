package main

import (
	"fmt"
	"strings"
)

type CurrencyMap = map[string]map[string]float64

func main() {

	countValue, startCurrency, targetCurreny := getInputData()
	totalConvertValue := convertValueByValuteName(countValue, startCurrency, targetCurreny)

	fmt.Println("----------Перевод из:", startCurrency, "в", targetCurreny, "=", totalConvertValue, "----------")

}

func convertValueByValuteName(value float64, curr string, target string) float64 {

	mapRation := CurrencyMap{
		"eur": {"usd": 1.14, "rub": 90.5},
		"usd": {"eur": 0.88, "rub": 79.25},
		"rub": {"usd": 0.013, "eur": 0.011},
	}

	return mapRation[curr][target] * value
}

func getInputData() (float64, string, string) {

	countValue := getTotalValue()

	fmt.Println("Введите стартовую валюту USD/EUR/RUB")
	startCurrency := getCurrencyName("")

	fmt.Println("Введите целевую валюту USD/EUR/RUB")
	targetCurreny := getCurrencyName(startCurrency)

	return countValue, startCurrency, targetCurreny

}

func getCurrencyName(startCurrency string) string {
	var currency string

	for {
		fmt.Scanln(&currency)

		isUsd := strings.ToLower(currency) == "usd"
		isEur := strings.ToLower(currency) == "eur"
		isRub := strings.ToLower(currency) == "rub"

		if !(isUsd || isEur || isRub) {
			fmt.Println("Введите корректную валюту из перечисленных")
			continue
		}

		if strings.EqualFold(currency, startCurrency) {
			fmt.Println("перевод в одну и ту же валюту не возможен, введите другую валюту")
			continue
		}

		fmt.Println("---Вы успешно ввели----")
		return strings.ToLower(currency)
	}

}

func getTotalValue() float64 {
	var countValue float64
	fmt.Println("Введите кол-во валюты больше нуля")

	for {
		_, err := fmt.Scanln(&countValue)

		if err != nil || countValue <= 0 {
			var discard string
			fmt.Scanln(&discard)
			fmt.Println("некорректный ввод, введите еще раз ")
			continue
		}
		fmt.Println("---Вы успешно ввели----")
		return countValue
	}
}
