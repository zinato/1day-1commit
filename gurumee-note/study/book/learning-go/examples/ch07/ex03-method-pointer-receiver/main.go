package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total += 1
	c.lastUpdated = time.Now()
}

func (c Counter) WrongIncrement() {
	c.total += 1
	c.lastUpdated = time.Now()
}

func (c Counter) ToString() string {
	return fmt.Sprintf("total: %v, last updated: %v", c.total, c.lastUpdated)
}
func main() {
	c := Counter{}
	fmt.Println(c.ToString())

	c.Increment()
	fmt.Println(c.ToString())

	c.WrongIncrement()
	fmt.Println(c.ToString())
}
