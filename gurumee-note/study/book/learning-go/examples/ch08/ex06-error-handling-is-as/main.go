package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

type MyErr struct {
	Codes []int
}

var myError = MyErr{
	Codes: []int{404},
}

func (me MyErr) Error() string {
	return fmt.Sprintf("codes: %v", me.Codes)
}

func (me MyErr) Is(target error) bool {
	if me2, ok := target.(MyErr); ok {
		return reflect.DeepEqual(me, me2)
	}

	return false
}

func fileChecker(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return myError
	}
	f.Close()
	return nil
}

func main() {
	err := fileChecker("not_here.txt")
	var myErr MyErr

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("that file doesn't exist")
		}

		if errors.Is(err, myError) {
			fmt.Println("is my error")
		}

		if errors.As(err, &myErr) {
			fmt.Println(myErr.Codes)
		}
	}
}
