package main

import (
	"encoding/csv"
	"fmt"
	"learning-go/examples/ch14/ex03-csv-marshalling/marshal"
	"strings"
)

type MyData struct {
	Name   string `csv:"name"`
	Age    string `csv:"age"`
	HasPet string `csv:"has_pet"`
}

func main() {
	data := `name,age,has_pet
"Jon","100",true
"Fred","42",false
"Martha","37",true
`
	r := csv.NewReader(strings.NewReader(data))
	allData, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	var entries []MyData
	marshal.UnMarshal(allData, &entries)
	fmt.Println(entries)

	out, err := marshal.Marshal(entries)
	if err != nil {
		panic(err)
	}

	sb := &strings.Builder{}
	w := csv.NewWriter(sb)
	w.WriteAll(out)
	fmt.Println(sb)
}
