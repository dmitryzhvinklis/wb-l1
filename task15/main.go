package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

// Задание 15
// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }
// func main() {
//   someFunc()
// }

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data 
	fmt.Println(hdr)

	justString = v[:100]

	hdr = (*reflect.SliceHeader)(unsafe.Pointer(&justString)).Data !
	fmt.Println(hdr)
}


func someFuncCorrect1() {
	v := createHugeString(1 << 10)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data
	fmt.Println(hdr)


	justString = strings.Clone(v[:100])
	hdr = (*reflect.SliceHeader)(unsafe.Pointer(&justString)).Data
	fmt.Println(hdr)
}


func someFuncCorrect2() {
	v := createHugeString(1 << 10)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data
	fmt.Println(hdr)

	justString = string([]rune(v[:100]))
	hdr = (*reflect.SliceHeader)(unsafe.Pointer(&justString)).Data
	fmt.Println(hdr)
}

func createHugeString(n int) string {
	s := strings.Repeat("a", n)
	return s
}

func main() {
	someFunc()
	fmt.Println()
	someFuncCorrect1()
	fmt.Println()
	someFuncCorrect2()
}
