package main

import "fmt"

func main() {
	words := make([]string, 0)
	words = append(words, "a", "cow", "smile", "gopher", "octopus", "anthropologist")
	for _, word := range words {
		switch size := len(word); {
		case size < 5:
			fmt.Println(word, "is a short word")
		case size == 5:
			fmt.Println(word, "is length ", size)
		case 5 < size && size < 10:
		default:
			fmt.Println(word, "is long word")
		}
	}
}
