package main

import "fmt"

type person struct {
	name string
	age  int
}

type pet struct {
	name string
	age  int
}

func main() {
	p1 := person{
		name: "gurumee",
		age:  31,
	}
	fmt.Println(p1)

	// 포인터
	pet1 := new(pet)
	pet1.name = "pet"
	pet1.age = 1
	fmt.Println(pet1)

	var pet2 struct {
		name string
		age  int
	}
	pet2 = p1
	fmt.Println(pet2)

	pet3 := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "cat",
	}
	fmt.Println(pet3)
}
