package main

// Задание 17
// Реализовать бинарный поиск встроенными методами языка.

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)


func binarySearch(v []int, val int) int {
	if !sort.IntsAreSorted(v) {
		panic("slice is not sorted")
	}
	offset := 0
	for len(v) != 0 {
		mid := len(v) / 2
		switch {
		case val == v[mid]:
			return offset + mid
		case val < v[mid]:
			v = v[:mid]
		case val > v[mid]:
			v = v[mid+1:]
			offset += mid + 1
		}
	}
	return -1
}


func main() {
	const sliceSize = 10000 
	const maxStep = 50      
	rand.Seed(time.Now().UnixNano())

	firstElement := rand.Intn(sliceSize/2) - rand.Intn(sliceSize) 


	randStep := func() int {
		if maxStep == 1 {
			return 1
		}
		return rand.Intn(maxStep-1) + 1 
	}

	testSlice := make([]int, sliceSize) 
	notInSlice := []int{firstElement - randStep()}


	for i, k := 0, firstElement; i < sliceSize; i, k = i+1, k+randStep() {
		testSlice[i] = k
		if i == 0 {
			continue
		}
		for n := testSlice[i-1] + 1; n < k; n++ {
			notInSlice = append(notInSlice, n)
		}
	}
	notInSlice = append(notInSlice, testSlice[sliceSize-1]+randStep())
	errCount := 0


	for _, val := range testSlice {
		idx := binarySearch(testSlice, val)
		if testSlice[idx] != val { 
			errCount++
			fmt.Printf("idx=%d, val=%d, testSlice[%d]=%d\n", idx, val, idx, testSlice[idx])
		}
	}


	for _, val := range notInSlice {
		if idx := binarySearch(testSlice, val); idx != -1 {
			fmt.Printf("val=%d, idx=%d (must be -1), testSlice[%d]=%d\n", val, idx, idx, testSlice[idx])
			errCount++
		}
	}
	if errCount > 0 {
		fmt.Printf("%d errors ocuried\n", errCount)
		os.Exit(1)
	}
	fmt.Println("OK")
}
