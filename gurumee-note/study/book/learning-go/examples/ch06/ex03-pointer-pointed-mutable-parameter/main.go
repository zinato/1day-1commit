package main

import "fmt"

func updateFailedPointer(p *int) {
	x := 10
	p = &x
}

func updatePointer(p *int) {
	*p *= 2
}

func main() {
	var p *int
	fmt.Printf("addr: %v\n", p)

	updateFailedPointer(p)
	fmt.Printf("addr: %v\n", p)

	x := 5
	fmt.Printf("value: %v, addr: %v\n", x, &x)

	updateFailedPointer(&x)
	fmt.Printf("value: %v, addr: %v\n", x, &x)

	updatePointer(&x)
	fmt.Printf("value: %v, addr: %v\n", x, &x)
}
