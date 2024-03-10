package main

// Задание 8
// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

import "fmt"


func setBit(num uint64, i int, set bool) uint64 {
	if set {
		num = num | (1 << i) 
	} else {
		num = num ^ (1 << i) 
	}
	return num
}

func main() {
	var n uint64
	n = setBit(n, 10, true)
	fmt.Println(n)
	n = setBit(n, 2, true) 
	fmt.Println(n)
	n = setBit(n, 10, false) 
	fmt.Println(n)
	n = setBit(n, 0, true) 
	fmt.Println(n)         
}
