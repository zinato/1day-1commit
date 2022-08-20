package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	for i, v := range arr {
		if i == 0 {
			continue
		}

		if i == len(arr)-1 {
			break
		}
		fmt.Println(i, v)
	}

	for i := 1; i < len(arr)-1; i++ {
		fmt.Println(i, arr[i])
	}

	m := make(map[string]string)
	m["hello"] = "world"
	m["go"] = "lang"

	for k, v := range m {
		fmt.Println(k, v)
	}
}
