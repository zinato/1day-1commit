package main

import "fmt"

type person struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

func generateStringPointer(s string) *string {
	return &s
}

func main() {
	p := person{
		FirstName:  "Pat",
		MiddleName: generateStringPointer("Perry"),
		LastName:   "Perterson",
	}
	fmt.Println(p)
}
