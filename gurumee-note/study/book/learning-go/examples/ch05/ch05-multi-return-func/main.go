package main

import (
	"errors"
	"fmt"
	"os"
)

func divAndMod(n, d int) (int, int, error) {
	if d == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	return n / d, n % d, nil
}

func main() {
	d, m, err := divAndMod(5, 3)
	fmt.Println(d, m, err)

	d, m, err = divAndMod(3, 0)
	fmt.Println(d, m, err)

	_, _, err = divAndMod(5, 0)
	d, m, _ = divAndMod(5, 2)
	fmt.Println(d, m)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
