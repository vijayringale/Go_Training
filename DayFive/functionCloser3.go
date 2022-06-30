package main

import "fmt"

func test2(myFunc func(int) int) {
	fmt.Println(myFunc(3))
}
	var counter int
func increment(i int) {
i = i + 1
}

func increment1(i *int) {
	*i = *i + 1
	}

func main( ){
	test := func(x int) int {
		return x* 4
	}
	test2(test)
	counter = 0
	increment(counter)
	fmt.Printf("%v", counter) 
	counter = 6
	increment1(&counter)
	fmt.Printf("%v", counter)
}