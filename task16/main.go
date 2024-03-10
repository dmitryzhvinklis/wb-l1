package main

// Задание 16
// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

import (
	"fmt"
	"math/rand"
	"time"
)


func quickSortRecursive(arr []int) {
	if len(arr) < 2 {
		return
	}

	p := partition(arr)
	if p > 1 {
		quickSortRecursive(arr[:p]) 
	}
	if p < len(arr)-2 {
		quickSortRecursive(arr[p+1:])
	}
}


func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	type sortRange struct {
		left, right int
	}
	sortStack := []sortRange{{left: 0, right: len(arr)}} 
	for len(sortStack) != 0 {
		sp := len(sortStack) - 1 
		
		sRange := sortStack[sp]
		sortStack = sortStack[:sp]

		
		p := partition(arr[sRange.left:sRange.right])
		
		if p > 1 {
			sortStack = append(sortStack, sortRange{left: sRange.left, right: sRange.left + p}) 
		}
		if sRange.right-(sRange.left+p+1) > 1 {
			sortStack = append(sortStack, sortRange{left: sRange.left + p + 1, right: sRange.right}) 
	}
}


func partition(arr []int) int {
    pivot := arr[len(arr)/2]
    left, right := 0, len(arr)-1
    for {
        for ; left < len(arr) && arr[left] < pivot; left++ {
        }
        for ; right >= 0 && arr[right] > pivot; right-- {
        }
        if left >= right {
            break
        }
        arr[left], arr[right] = arr[right], arr[left]
        left++
        right--
    }
    return right
}


func main() {
	const numOfTests = 1000
	const arrayLength = 10000
	rand.Seed(time.Now().UnixNano())
	fillFn := func() int {
		return rand.Intn(1000) - 500
	}
	
	fmt.Print("Testing quickSortRecursive... ")
	testSort(numOfTests, arrayLength, quickSortRecursive, fillFn)

	
	fmt.Print("Testing quickSort... ")
	testSort(numOfTests, arrayLength, quickSort, fillFn)
}


func testSort(numOfTests, arrayLength int, sortFn func([]int), fillFn func() int) {
	for i := 0; i < numOfTests; i++ {
		
		arr := make([]int, arrayLength)
		for idx := range arr {
			arr[idx] = fillFn()
		}
		backupArr := make([]int, len(arr))
		copy(backupArr, arr)
		quickSort(arr)
		if !isSorted(arr) {
			fmt.Printf("FAIL\n%v ---> %v\n", backupArr, arr)
		}
	}
	fmt.Println("OK")
}


func isSorted(arr []int) bool {
	if len(arr) < 2 {
		return true
	}
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}
