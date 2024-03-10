package main

// Задание 7
// Реализовать конкурентную запись данных в map.

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)




type threadSafeMapper interface {

	Store(key, value any)

	Load(key any) (any, bool)

	Delete(key any)
}

var _ threadSafeMapper = (*rwMutexMap)(nil)


type rwMutexMap struct {
	mu *sync.RWMutex
	m  map[any]any
}


func (rwm *rwMutexMap) Store(key, value any) {
	rwm.mu.Lock()
	defer rwm.mu.Unlock()
	rwm.m[key] = value
}


func (rwm *rwMutexMap) Load(key any) (any, bool) {
	rwm.mu.RLock()
	defer rwm.mu.RUnlock()
	val, ok := rwm.m[key]
	return val, ok
}


func (rwm *rwMutexMap) Delete(key any) {
	rwm.mu.Lock()
	defer rwm.mu.Unlock()
	delete(rwm.m, key)
}


var _ threadSafeMapper = (*sync.Map)(nil)


func testOneStorerManyReaders(sm threadSafeMapper) time.Duration {
	const sizeOfTestData = 1000000 
	const numReadWorkers = 200     
	const sizeOfMap = 1000000      
	const storerDelay = 1000       
	ch := make(chan int, 100)
	wg := &sync.WaitGroup{}


	storer := func() {
		for i := 0; i < sizeOfMap; i++ {
			val := rand.Int()
			sm.Store(i, val)

			time.Sleep(time.Microsecond * time.Duration(rand.Intn(storerDelay)))
		}
	}


	reader := func() {
		for key := range ch {
			sm.Load(key)
		}
		wg.Done()
	}
	start := time.Now()

	go storer()

	for i := 0; i < numReadWorkers; i++ {
		wg.Add(1)
		go reader()
	}
	
	for i := 0; i < sizeOfTestData; i++ {
		ch <- rand.Intn(sizeOfMap)
	}

	close(ch)

	wg.Wait()
	return time.Since(start)
}

func main() {
	rwm := &rwMutexMap{
		mu: &sync.RWMutex{},
		m:  make(map[any]any),
	}
	sm := &sync.Map{}

	fmt.Print("Running testOneStorerManyReaders for rwMutexMap...")
	fmt.Printf(" done in %v\n", testOneStorerManyReaders(rwm))

	fmt.Print("Running testOneStorerManyReaders for sync.Map...")
	fmt.Printf(" done in %v\n", testOneStorerManyReaders(sm))
}
