package main

import (
	"fmt"
	"log"
)

type person struct {
	name string
	age  int
}

func DontMakePerson(p *person, name string, age int) error {
	p.name = name
	p.age = age
	return nil
}

func MakePerson(name string, age int) (person, error) {
	p := person{name: name, age: age}
	return p, nil
}

func main() {
	p1 := person{}
	err := DontMakePerson(&p1, "perter", 20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(p1)

	p2, err := MakePerson("john", 15)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(p2)
}
