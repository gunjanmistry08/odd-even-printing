package main

func print_routine(start_number int, firstTurn, secondTurn, done chan struct{}) {
	for i := start_number; i <= 100; i = i + 2 {
		<-firstTurn

		println(i)
		if i != 100 {
			secondTurn <- struct{}{}
		}

	}
	done <- struct{}{}
}

func main() {

	even_turn := make(chan struct{})
	odd_turn := make(chan struct{})

	done := make(chan struct{})

	go print_routine(0, even_turn, odd_turn, done)
	go print_routine(1, odd_turn, even_turn, done)

	even_turn <- struct{}{}

	<-done
	<-done

}
