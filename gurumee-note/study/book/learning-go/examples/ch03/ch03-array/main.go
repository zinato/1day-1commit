package main

import "fmt"

func main() {
	var x = [3]int{10, 20, 30}
	fmt.Println(x)

	y := [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(y)

	x1 := [...]int{1, 2, 3}
	x2 := [3]int{1, 2, 3}
	fmt.Println(x1 == x2)
}
