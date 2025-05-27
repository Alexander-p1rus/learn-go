package main

import (
	"fmt"
)

func main() {
	const ratioUsdToEur = 0.88
	const ratioUsdToRub = 80.25

	ratioEurToRub := ratioUsdToRub / ratioUsdToEur

	fmt.Println(ratioEurToRub)
}
