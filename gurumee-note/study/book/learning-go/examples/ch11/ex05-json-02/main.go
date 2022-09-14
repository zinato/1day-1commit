package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	toFile := Person{
		Name: "Fred",
		Age:  40,
	}
	ifErrorExistLog := func(err error) {
		if err != nil {
			log.Fatalln(err)
		}
	}
	tmpFile, err := ioutil.TempFile(os.TempDir(), "sample-")
	ifErrorExistLog(err)

	defer func(name string) {
		err := os.Remove(name)
		ifErrorExistLog(err)
	}(tmpFile.Name())

	err = json.NewEncoder(tmpFile).Encode(toFile)
	ifErrorExistLog(err)

	err = tmpFile.Close()
	ifErrorExistLog(err)

	tmpFile2, err := os.Open(tmpFile.Name())
	ifErrorExistLog(err)

	var fromFile Person
	err = json.NewDecoder(tmpFile2).Decode(&fromFile)
	ifErrorExistLog(err)
	err = tmpFile2.Close()
	ifErrorExistLog(err)
	fmt.Printf("%+v\n", fromFile)
}
