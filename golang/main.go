package main

import (
	"sync"
)

func print_routine(start_number int, first_lock, second_lock *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start_number; i < 100; i = i + 2 {
		second_lock.Lock()
		println(i)
		first_lock.Unlock()
	}
}

func main() {
	var odd_lock = sync.Mutex{}
	var even_lock = sync.Mutex{}
	var wg sync.WaitGroup
	odd_lock.Lock()
	wg.Add(2)
	go print_routine(0, &odd_lock, &even_lock, &wg)
	go print_routine(1, &even_lock, &odd_lock, &wg)
	wg.Wait()
}
