package main

import (
	"fmt"
)

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}

func main() {
	arr := []int{1, 2, 0, 5}
	for _, e := range arr {
		div60(e)
	}
}
