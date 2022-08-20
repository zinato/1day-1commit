package main

import (
	"fmt"
)

func main() {
	x := []int{10, 20, 30}
	fmt.Println(len(x))

	var x2 []int
	fmt.Println(x2 == nil)

	x2 = append(x2, 10)
	fmt.Println(x2, len(x2), cap(x2))
	x2 = append(x2, 20, 30, 40)
	fmt.Println(x2, len(x2), cap(x2))

	x3 := make([]int, 5, 10)
	x3 = append(x3, 1)
	x2 = append(x2, x3...)
	fmt.Println(x2, len(x2), cap(x2))
	x2 = append(x2, 20, 30, 40)
	fmt.Println(x2, len(x2), cap(x2))

	y := x2[2:5]
	z := x2[2:5:5]
	fmt.Printf("x2: %v, y: %v, z: %v\n", x2, y, z)

	y[1] = 5
	z[1] = 6
	fmt.Printf("x2: %v, y: %v, z: %v\n", x2, y, z)

	y = append(y, 69)
	z = append(z, 699)
	fmt.Printf("x2: %v, y: %v, z: %v\n", x2, y, z)
	// "full slice expression"으로 append 되면 그 때부터 다른 슬라이스가 된다.
	y[2] = 5
	z[2] = 6
	fmt.Printf("x2: %v, y: %v, z: %v\n", x2, y, z)

	x4 := make([]int, 10)
	copy(x4, x2)
	fmt.Println(x4)
}
