package main

import (
	"fmt"
	"math/bits"
	"unsafe"
)

var isLE bool

func init() {
	var x uint16 = 0xFF00
	xb := *(*[2]byte)(unsafe.Pointer(&x))
	isLE = (xb[0] == 0x00)
}

type Data struct {
	Value  uint32
	Label  [10]byte
	Active bool
}

func DataFromBytesUnsafe(b [16]byte) Data {
	data := *(*Data)(unsafe.Pointer(&b))
	if isLE {
		data.Value = bits.ReverseBytes32(data.Value)
	}

	return data
}

func BytesFromDataUnsafe(d Data) [16]byte {
	if isLE {
		d.Value = bits.ReverseBytes32(d.Value)
	}

	b := *(*[16]byte)(unsafe.Pointer(&d))
	return b
}

func main() {
	d := Data{
		Value:  3,
		Label:  [10]byte{0},
		Active: true,
	}

	b := BytesFromDataUnsafe(d)
	fmt.Println(b)

	d2 := DataFromBytesUnsafe(b)
	fmt.Println(d2)
}
