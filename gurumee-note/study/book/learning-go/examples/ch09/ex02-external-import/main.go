package main

import (
	"fmt"
	"log"

	"github.com/shopspring/decimal"
)

func main() {
	amount, err := decimal.NewFromString("50")
	if err != nil {
		log.Fatalln(err)
	}

	percent, err := decimal.NewFromString("20")
	if err != nil {
		log.Fatalln(err)
	}

	percent = percent.Add(decimal.NewFromInt(100))
	total := amount.Add(amount.Mul(percent).Round(2))
	fmt.Println(amount, percent, total.StringFixed(2))
}
