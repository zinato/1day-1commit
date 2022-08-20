package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	s := `Hello 
"World"`
	fmt.Println(s)

	var flag bool
	var isAwesome = true
	fmt.Println(flag, isAwesome)

	var x int = 10
	x *= 2
	fmt.Println(x)

	c1 := complex(2.5, 3.1)
	c2 := complex(10.2, 2)
	fmt.Println(c1 + c2)
	fmt.Println(c1 - c2)
	fmt.Println(c1 * c2)
	fmt.Println(c1 / c2)
	fmt.Println(real(c1))
	fmt.Println(imag(c1))
	fmt.Println(cmplx.Abs(c1))

	var y float64 = 30.2
	var z float64 = float64(x) + y
	var d int = x + int(z)
	fmt.Println(d)
}
