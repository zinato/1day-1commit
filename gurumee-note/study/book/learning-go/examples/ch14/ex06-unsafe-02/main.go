package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func main() {
	s := "hello"
	sHdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println(sHdr.Len)

	for i := 0; i < sHdr.Len; i++ {
		bp := *(*byte)(unsafe.Pointer(sHdr.Data + uintptr(i)))
		fmt.Print(string(bp))
	}
	fmt.Println()
	runtime.KeepAlive(s)

	a := []int{10, 20, 30}
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	fmt.Println(aHdr.Len)
	fmt.Println(aHdr.Cap)

	intByteSize := unsafe.Sizeof(a[0])
	fmt.Println(intByteSize)
	for i := 0; i < aHdr.Len; i++ {
		intVal := *(*int)(unsafe.Pointer(aHdr.Data + intByteSize*uintptr(i)))
		fmt.Println(intVal)
	}
	runtime.KeepAlive(a)
}
