package main

import (
	"fmt"
	"os"
)

type MyInt int

func main() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	i2, ok := i.(MyInt)
	if !ok {
		_ = fmt.Errorf("unexpected type for %v", i)
		os.Exit(1)
	}
	fmt.Println(i2 + 1)

	switch i.(type) {
	case nil:
		fmt.Println("It is nil")
	case int:
		fmt.Println("It is int")
	case MyInt:
		fmt.Println("It is MyInt")
	default:
		fmt.Println("fuck")
	}

	i3, ok := i.(int)
	if !ok {
		fmt.Printf("unexpected type for %v\n", i)
		os.Exit(1)
	}
	fmt.Println(i3 + 10)

}
