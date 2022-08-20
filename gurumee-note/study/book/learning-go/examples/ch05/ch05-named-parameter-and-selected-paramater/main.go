package main

import "fmt"

type MyFuncOpt struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpt) error {
	fmt.Println(opts)
	return nil
}

func main() {
	MyFunc(MyFuncOpt{
		LastName: "Patel",
		Age:      50,
	})
	MyFunc(MyFuncOpt{
		FirstName: "Joe",
		LastName:  "Smith",
	})
}
