package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x, &x)

	if x > 5 {
		fmt.Println(x, &x)
		x, y := 5, 20
		fmt.Println(x, &x)
		fmt.Println(y, &y)
	}

	fmt.Println(x, &x)
}
