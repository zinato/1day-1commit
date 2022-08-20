package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["test"] = 1
	m["go"] = 2
	fmt.Println(m)

	m["test"] = 3
	fmt.Println(m)

	v, ok := m["test"]
	fmt.Println(v, ok)

	v, ok = m["hello"]
	fmt.Println(v, ok)

	delete(m, "go")
	fmt.Println(m)

	m["test2"] = 5

	for k, v := range m {
		fmt.Println(k, v)
	}
}
