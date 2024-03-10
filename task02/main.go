package main

// Задание 2
// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых
// из массива (2,4,6,8,10) и выведет их квадраты в stdout.

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	wg := new(sync.WaitGroup)
	for _, num := range numbers {
		wg.Add(1)

		go func(n int) {
			fmt.Println(n * n)
			wg.Done()
		}(num)
	}
	// ждём окончания
	wg.Wait()
}
