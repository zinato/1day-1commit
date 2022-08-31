package main

import (
	"fmt"
	"log"
	"os"
)

func fileChecker(fileName string) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in fileChecker defer: %w", err)
		}
	}()

	f, err := os.Open(fileName)
	if err != nil {
		return err
	}

	f.Close()
	return nil
}

func main() {
	fileName := "not_here.txt"
	err := fileChecker(fileName)
	if err != nil {
		log.Fatalln(err)
	}
}
