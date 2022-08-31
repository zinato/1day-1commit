package main

import (
	"errors"
	"fmt"
	"log"
)

func DivAndMod(n, d int) (int, int, error) {
	if d == 0 {
		return 0, 0, errors.New("denominator is 0")
	}

	return n / d, n % d, nil
}

func main() {
	n, d := 20, 3
	div, mod, err := DivAndMod(n, d)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(div, mod)
}
