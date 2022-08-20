package main

import (
	"fmt"
	"strconv"
)

func add(i, j int) int {
	return i + j
}

func sub(i, j int) int {
	return i - j
}
func mul(i, j int) int {
	return i * j
}
func div(i, j int) int {
	return i / j
}

type opFuncType func(int, int) int

var opMap = map[string]opFuncType{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, expr := range expressions {
		if len(expr) != 3 {
			fmt.Println("invalid expression", expr)
			continue
		}

		p1, err := strconv.Atoi(expr[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		op := expr[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator", op)
			continue
		}

		p2, err := strconv.Atoi(expr[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		result := opFunc(p1, p2)
		fmt.Println(result)
	}
}
