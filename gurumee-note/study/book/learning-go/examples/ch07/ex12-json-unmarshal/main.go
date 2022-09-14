package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

const JSON = `{
    "a": "b"
    "b": 4,
	"c": true,
}`

func main() {
	data := map[string]interface{}{}
	fileName := "/Users/gurumee92/Studies/1day-1commit/gurumee-note/study/book/learning-go/examples/ch07/ex12-json-unmarshal/sample.json"
	c, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(c, &data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)
}
