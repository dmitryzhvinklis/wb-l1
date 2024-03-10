package main

// Задание 14
// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

import "fmt"


func whatsThis(x interface{}) {
	fmt.Printf("x (%T) = %+v\n", x, x)
}


func xOperation(x interface{}) interface{} {
	switch y := x.(type) {

	case int:
		return y * y

	case string:
		runes := []rune(y)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)

	case bool:
		return !y

	case chan string:
		go func() {
			y <- "Hello!"
			close(y)
		}()
		return nil

	case []int:
		var sum int
		for _, n := range y {
			sum += n
		}
		return sum
	}
	return nil
}

func main() {
	xs := []interface{}{123456789, "Preved medved!", true, make(chan string), []int{1, -2, 3, -4, 5, -6, 7}}
	for _, x := range xs {
		whatsThis(x)
		if chX, ok := x.(chan string); ok {
			xOperation(x)
			fmt.Printf("xOperation(x) sended %q to the channel\n\n", <-chX)
			continue
		}
		fmt.Printf("xOperation(x) = %v\n\n", xOperation(x))
	}
}
