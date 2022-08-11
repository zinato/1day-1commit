package main

import "fmt"

func main() {
	s := "Hello ☃︎"
	s2 := s[4:7]
	s3 := s[4:]
	fmt.Printf("s2: %v, s3: %v\n", s2, s3)

	var r rune = 'x'
	s4 := string(r)

	var b byte = 'y'
	s5 := string(b)
	fmt.Println(s4, s5)

	s6 := "Hello"
	bs := []byte(s6)
	rs := []rune(s6)
	fmt.Println(bs, rs)
}
