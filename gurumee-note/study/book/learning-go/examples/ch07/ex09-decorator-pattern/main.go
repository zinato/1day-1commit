package main

import "fmt"

type LoginProvider struct {
}

func (lp LoginProvider) Process(data string) string {
	return data + " process!"
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	data := "Hello World"
	processedData := c.L.Process(data)
	fmt.Println(processedData)
}

func main() {
	c := Client{
		L: LoginProvider{},
	}
	c.Program()
}
