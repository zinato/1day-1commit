package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	ch := make(chan int, len(a))

	for _, v := range a {
		v := v
		go func() {
			ch <- v * 2
		}()
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}
