package main

import (
	"fmt"
	"learning-go/math"
	"learning-go/print"
)

func main() {
	n := math.Double(2)
	output := print.Format(n)
	fmt.Println(output)
}
