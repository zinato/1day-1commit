package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func process(r io.Reader) error {
	data := make([]byte, 100)

	for {
		count, err := r.Read(data)
		if err != nil {
			return err
		}

		if count == 0 {
			return nil
		}
		fmt.Println(data)
	}
}

func main() {
	fileName := "/Users/gurumee92/Studies/1day-1commit/gurumee-note/study/book/learning-go/examples/ch07/ex09-duck-typing/main.go"
	r, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	defer r.Close()
	err = process(r)
	if err != nil {
		log.Fatalln(err)
	}
}
