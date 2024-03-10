package main

// Задание 23
// Удалить i-ый элемент из слайса.

import "fmt"


type slice[T any] []T

func (s *slice[T]) deleteElement(i int) {
	if i < 0 || i > len(*s) {
		panic("index out of range")
	}
	
	*s = append((*s)[:i], (*s)[i+1:]...)
}


func main() {

	n := slice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	n.deleteElement(5)
	fmt.Println(n)


	s := slice[string]{"zero", "one", "two", "three", "four", "five"}
	s.deleteElement(3)
	fmt.Println(s)
}
