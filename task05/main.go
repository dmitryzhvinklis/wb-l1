package main

// Задание 5
// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

import (
	"context"
	"fmt"
	"sync"
	"time"
)


func writer(ctx context.Context, ch chan<- int) {
	for i := 0; true; i++ { 
		select {
		default:
			ch <- i
			time.Sleep(time.Millisecond * 10) 
		case <-ctx.Done():
			close(ch) 
			return
		}
	}
}


func reader(ch <-chan int, wg *sync.WaitGroup) {
	for num := range ch {
		fmt.Printf("%d\t", num) 
	}
	fmt.Println()
	wg.Done() 
}

func main() {
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()            
	wg := new(sync.WaitGroup) 
	ch := make(chan int, 1)
	wg.Add(1)
	go writer(ctx, ch)
	go reader(ch, wg)
	wg.Wait() 
}
