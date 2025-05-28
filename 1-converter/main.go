package main

import (
	"fmt"
	"strings"
)

func main() {

	countValue, startCurrency, targetCurreny := getInputData()
	totalConvertValue := convertValueByValuteName(countValue, startCurrency, targetCurreny)

	fmt.Println("----------Перевод из:", startCurrency, "в", targetCurreny, "=", totalConvertValue, "----------")

}

func convertValueByValuteName(value float64, curr string, target string) float64 {

	const UsdToEurRatio = 0.88
	const UsdToRubRatio = 80.25
	const RubToEurRatio = 0.011

	switch curr {
	case "rub":
		if target == "eur" {
			return value * RubToEurRatio
		} else if target == "usd" {
			return value / UsdToRubRatio
		}

	case "eur":
		if target == "usd" {
			return value / UsdToEurRatio
		} else if target == "rub" {
			return value / RubToEurRatio
		}
	case "usd":
		if target == "rub" {
			return value * UsdToRubRatio
		} else if target == "eur" {
			return value * UsdToEurRatio
		}
	}

	return 1337
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
			fmt.Println(!isUsd || !isEur)
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
