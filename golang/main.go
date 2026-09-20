package main

import "time"

func printRoutine(startNumber int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := startNumber; ; i += 2 {
			out <- i
		}
	}()

	return out
}

func main() {

	even_chan := printRoutine(0)
	odd_chan := printRoutine(1)

	for i := 0; i < 100; i++ {
		select {
		case num1 := <-even_chan:
			println(num1)
		case num2 := <-odd_chan:
			println(num2)
		}
	}

	time.Sleep(time.Second * 3)

}
