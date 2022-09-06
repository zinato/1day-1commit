package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Order struct {
	ID          string    `json:"id"`
	DateOrdered time.Time `json:"date_ordered"`
	CustomerID  string    `json:"customer_id"`
	Items       []Item    `json:"items"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	data := `
{
	"id": "1235",
	"date_ordered": "2020-05-01T13:01:02Z",
	"customer_id": "3",
	"items": [{"id": "xy12", "name": "Thing1"},{"id": "xy13", "name": "Thing2"}]
}`
	var o Order
	err := json.Unmarshal([]byte(data), &o)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(o)

	out, err := json.Marshal(o)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(out))
}
