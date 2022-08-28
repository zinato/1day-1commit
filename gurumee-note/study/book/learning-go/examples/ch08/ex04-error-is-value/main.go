package main

import "fmt"

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func GenerateError(flag bool) error {
	if flag {
		return StatusErr{
			Status:  NotFound,
			Message: "Not Found",
		}
	}
	return nil
}

func main() {
	err := GenerateError(true)
	fmt.Println(err, err != nil)

	err = GenerateError(false)
	fmt.Println(err, err != nil)
}
