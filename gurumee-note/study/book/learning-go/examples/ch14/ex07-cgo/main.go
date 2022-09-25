package main

import "fmt"

/*
	int add(int a, int b){
		return a + b;
	}

	int mul(int a, int b) {
		return a * b;
	}
*/
import "C"

func main() {
	add := C.add(5, 2)
	fmt.Println(add)

	mul := C.mul(5, 2)
	fmt.Println(mul)
}
