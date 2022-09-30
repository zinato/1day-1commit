package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A int    `myTag:"A"`
	B string `myTag:"B"`
}

func main() {
	var x int
	xt := reflect.TypeOf(x)
	fmt.Println(xt.Name())

	xpt := reflect.TypeOf(&x)
	fmt.Println(xpt.Name())
	fmt.Println(xpt.Kind())
	fmt.Println(xpt.Elem().Name())
	fmt.Println(xpt.Elem().Kind())

	f := Foo{}
	ft := reflect.TypeOf(f)
	for i := 0; i < ft.NumField(); i++ {
		cur := ft.Field(i)
		fmt.Println(cur.Name, cur.Type.Name())

		tag := cur.Tag.Get("myTag")
		fmt.Println(tag, reflect.ValueOf(tag))
	}

	i := 10
	iv := reflect.ValueOf(&i)
	ivv := iv.Elem()
	ivv.SetInt(20)
	fmt.Println(i)
}
