package main

import "fmt"

func main() {

	consume()
}
func consume() chan<- int {
	ch := ConsumeChannel()
	for i := 0; i < 10; i++ {
		ch <- i + 1
		fmt.Println("Produced", i+1)
	}
	return ch
}
func ConsumeChannel() chan<- int {

	ch1 := make(chan int)

	go func() {
		for i := range ch1 {
			fmt.Println("consumed", i)

		}
	}()

	return ch1
}
