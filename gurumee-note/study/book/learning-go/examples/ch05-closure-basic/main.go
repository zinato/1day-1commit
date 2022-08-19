package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	people := make([]Person, 0)
	people = append(people, Person{"Pat", "Patterson", 37})
	people = append(people, Person{"Tracy", "Bobbert", 23})
	people = append(people, Person{"Fred", "Fredson", 18})

	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
}
