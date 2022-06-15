package main

import "fmt"

func Arrsa_Addition(p int) {

	var array [100]int
	slice1 := array[1:20]
	for i := 0; i >= p; i++ {
		array = append(array, i)
	}
	slice1 := array[1:20]
	var sum = 0
	for _, i := range slice1 {
		sum = sum + i
		fmt.Println(sum)
	}

}

func main() {

	// Calling Goroutine
	go Arrsa_Addition(100)

	// Calling normal function
	Arrsa_Addition()
}
