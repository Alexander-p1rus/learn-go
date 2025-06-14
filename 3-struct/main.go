package main

import (
	"fmt"
	"time"

	"demo/app-1/3-struct/bins"
)

func main() {

	bin, err := bins.СreateBin("222", true, time.Now(), "hello")

	if err != nil {
		panic(err)
	}

	fmt.Printf("(*bin): %v\n", bin)

}
