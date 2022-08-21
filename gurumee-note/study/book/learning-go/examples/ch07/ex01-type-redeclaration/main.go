package main

import "fmt"

type Score int

func (s Score) GetScore() int {
	return int(s)
}

func main() {
	s := Score(5)
	fmt.Println(s.GetScore())
}
