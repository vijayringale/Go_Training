package main

import (
	"fmt"
)

type Point struct {
	a   int
	b   int
	arr []int
	sum int
}

func runtine1(a int, b int, arr []int, sum int) {
	slice := arr[a:b]

	for _, items := range slice {
		fmt.Print(items, " ")
		sum = sum + items
	}

	fmt.Print("\n\n", sum, "\n")
}

func main() {

	var array []int

	for i := 0; i < 100; i++ {
		array = append(array, i)
	}
	fmt.Println(array)

	go runtine1(0, 20, array, 0)
	runtine1(20, 40, array, 0)
	go runtine1(40, 60, array, 0)
	runtine1(60, 80, array, 0)
	runtine1(80, 99, array, 0)

}
