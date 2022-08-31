package main

import "fmt"

type Inner struct {
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %v", val)
}

func (i Inner) Double() string {
	return i.IntPrinter(i.A * 2)
}

type Outer struct {
	Inner
	S string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %v", val)
}

func main() {
	o := Outer{
		Inner: Inner{A: 10},
		S:     "Hello",
	}
	fmt.Println(o.Double())
}
