package main

import "fmt"

func xor(ch1 chan int) {
	for i := 0; i < 11; i++ {
		ch1 <- i + 1
		fmt.Println("Produced", i+1)
	}
	close(ch1)
}

func ant(ch2 chan int) {
	for i := 11; i < 21; i++ {
		ch2 <- i + 1
		fmt.Println("Produced", i+1)
	}
	close(ch2)
}
func Xoriant(ch1, ch2, ch3 chan int) {
	for v := range ch1 {
		ch3 <- v
	}
	for v1 := range ch2 {
		ch3 <- v1
	}
	close(ch3)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go xor(ch1)
	go ant(ch2)
	go Xoriant(ch1, ch2, ch3)
	for i := range ch3 {
		fmt.Println("Comsumed", i)
	}
}
