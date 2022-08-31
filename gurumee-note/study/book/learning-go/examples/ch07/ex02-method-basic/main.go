package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) ToString() string {
	return fmt.Sprintf("%v %v, age: %v", p.FirstName, p.LastName, p.Age)
}

func main() {
	p := Person{
		FirstName: "Fred",
		LastName:  "Fredson",
		Age:       52,
	}
	fmt.Println(p)
	fmt.Println(p.ToString())
}
