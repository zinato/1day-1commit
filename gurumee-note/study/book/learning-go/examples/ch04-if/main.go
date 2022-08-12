package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if n := rand.Intn(10); n == 0 {
		fmt.Println("zero", n)
	} else if n > 5 {
		fmt.Println("range 5 < x", n)
	} else {
		fmt.Println("range 0 < x <= 5", n)
	}
}
