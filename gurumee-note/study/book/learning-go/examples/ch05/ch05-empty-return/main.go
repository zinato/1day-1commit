package main

import (
	"errors"
	"fmt"
	"os"
)

func divAndMod(n, d int) (div, mod int, err error) {
	if d == 0 {
		err = errors.New("cannot divide zero")
		return
	}

	div, mod = n/d, n%d
	return
}

func main() {
	x, y, err := divAndMod(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(x, y)
}
