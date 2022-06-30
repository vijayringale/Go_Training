package main

import "fmt"

func main() {
	test := func(x int) int {
		return x * 2
	}(3)

	jony := test
	fmt.Println("Value Of Test Function := ", test ,jony,)
}