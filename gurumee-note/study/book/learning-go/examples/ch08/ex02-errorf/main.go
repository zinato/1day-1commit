package main

import (
	"fmt"
	"log"
)

func doubleEven(i int) (int, error) {
	if i%2 == 0 {
		return 0, fmt.Errorf("%v isn't an even number", i)
	}

	return i * 2, nil
}
func main() {
	i := 5
	res, err := doubleEven(i)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
	i = 6
	res, err = doubleEven(i)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(i)
}
