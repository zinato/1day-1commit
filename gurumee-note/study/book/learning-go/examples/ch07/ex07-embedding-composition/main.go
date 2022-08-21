package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) ToString() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindEmployees(name string) bool {
	for _, e := range m.Reports {
		if name == e.Name {
			return true
		}
	}

	return false
}

func main() {
	e1 := Employee{Name: "Bob", ID: "11111"}
	e2 := Employee{Name: "Alice", ID: "11112"}

	m := Manager{
		Employee: Employee{
			Name: "Joe",
			ID:   "01111",
		},
		Reports: []Employee{e1, e2},
	}

	fmt.Println(m.ToString())
	fmt.Println(m.FindEmployees("Bob"))
	fmt.Println(m.FindEmployees("Alice"))
	fmt.Println(m.FindEmployees("Joe"))
}
