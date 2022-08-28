package main

import (
	"errors"
	"fmt"
	"os"
)

type Status int

const (
	InvalidLogin = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
	err     error
}

func (se StatusErr) Error() string {
	return se.Message
}

func (se StatusErr) Unwrap() error {
	return se.err
}

func fileChecker(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return StatusErr{
			Status:  NotFound,
			Message: "No exist files",
			err:     err,
		}
	}
	f.Close()
	return nil
}
func main() {
	err := fileChecker("not_here.txt")
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}

		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("this file doesn't exist")
		}
	}
}
