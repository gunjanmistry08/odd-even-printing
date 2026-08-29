package main

import (
	"sync"
)

// odd routine
func odd_routine(odd_lock, even_lock *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 100; i = i + 2 {
		odd_lock.Lock()
		println(i)
		even_lock.Unlock()
	}
}

// even routine
func even_routine(odd_lock, even_lock *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i = i + 2 {
		even_lock.Lock()
		println(i)
		odd_lock.Unlock()
	}
}
func main() {
	var odd_lock = sync.Mutex{}
	var even_lock = sync.Mutex{}
	var wg sync.WaitGroup
	odd_lock.Lock()
	wg.Add(2)
	go odd_routine(&odd_lock, &even_lock, &wg)
	go even_routine(&odd_lock, &even_lock, &wg)
	wg.Wait()
}
