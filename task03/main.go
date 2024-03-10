package main

// Задание 3
// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….)
// с использованием конкурентных вычислений.

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func useAtomic(numbers []int) int {
	var result int64

	wg := new(sync.WaitGroup)
	for _, num := range numbers {
		wg.Add(1)

		go func(n int) {

			atomic.AddInt64(&result, int64(n*n))
			wg.Done()
		}(num)
	}

	wg.Wait()
	return int(result)
}

func useChannels(numbers []int, chSize int) int {
	wg := new(sync.WaitGroup)

	chSqr := make(chan int, chSize)
	chResult := make(chan int)

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {

			chSqr <- n * n
			wg.Done()
		}(num)
	}

	go func() {
		var result int

		for num := range chSqr {
			result += num
		}
		chResult <- result
	}()

	wg.Wait()

	close(chSqr)
	return <-chResult
}

func main() {

	nums := []int{2, 4, 6, 8, 10}
	fmt.Println(useAtomic(nums), useChannels(nums, 5))

	// сравним скорость работы двух функций
	//
	// Результаты эксперимента показывают, что буфер канала ожидаемо становится узким местом,
	// т.к. горутинам приходится ждать возможности записать результат в канал.
	// Если размер буфера канала равен числу обрабатываемых элементов, то разницы в скорости
	// не обнаруживается. Но буфер канала забирает ощутимо больше памяти, поэтому подход с использованием
	// пакета atomic является более эффективным в данном случае.
	const sliceSize = 5000000
	const chanSize = 5000000
	numbers := make([]int, 0, sliceSize)
	for i := 1; i <= sliceSize; i++ {
		numbers = append(numbers, i*2)
	}
	start := time.Now()
	atomicRes := useAtomic(numbers)
	t1 := time.Since(start)
	start = time.Now()
	chanRes := useChannels(numbers, chanSize)
	t2 := time.Since(start)
	fmt.Printf("useAtomic:\tresult: %d, time: %v\n", atomicRes, t1)
	fmt.Printf("useChannels:\tresult: %d, time: %v\n", chanRes, t2)
}
