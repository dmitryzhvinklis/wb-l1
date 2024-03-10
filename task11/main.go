package main

// Задание 11
// Реализовать пересечение двух неупорядоченных множеств.

import "fmt"


func intersection[T comparable](setA, setB []T) []T {
	result := make([]T, 0, len(setA))
	mapB := make(map[T]bool)
	for _, item := range setB {
		mapB[item] = true
	}

	for _, item := range setA {
		if mapB[item] {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	

	
	fmt.Println(intersection(
		[]int{2, 6, 8, 4, 12, 14, 10, 18, 12, 14, 16},
		[]int{21, 18, 15, 9, 12, 3, 6},
	))


	fmt.Println(intersection(
		[]float32{-1.2, 5.6, 12.4, -12.4, -3.4},
		[]float32{12, 43, -22, 11, -1, 10},
	))

	
	fmt.Println(intersection(
		[]string{"рыба", "медуза", "кит", "акула", "краб", "кашалот", "осьминог", "каракатица", "кальмар", "дельфин", "креветка", "косатка"},
		[]string{"слон", "кошка", "дельфин", "собака", "кит", "хомяк", "волк", "кашалот", "корова", "козёл", "верблюд"},
	))
}
