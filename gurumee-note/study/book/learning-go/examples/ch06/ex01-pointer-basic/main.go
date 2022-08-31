package main

import (
	"fmt"
	"unsafe"
)

func main() {
	x := 10
	y := true
	pX := &x
	pY := &y
	fmt.Printf("value: %v, pointer: %v size: %v\n", x, &x, unsafe.Sizeof(x))
	fmt.Printf("value: %v, pointer: %v size: %v\n", y, &y, unsafe.Sizeof(y))
	fmt.Printf("value: %v, pointer: %v size: %v\n", *pX, pX, unsafe.Sizeof(pX))
	fmt.Printf("value: %v, pointer: %v size: %v\n", *pY, pY, unsafe.Sizeof(pY))
}
