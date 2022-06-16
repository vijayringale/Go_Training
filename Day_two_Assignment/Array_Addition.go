package main

import "fmt"

type Point struct {
	a   int
	b   int
	arr []int
	sum int
}

func runtine1(p Point) {
	slice := p.arr[p.a:p.b]

	for _, items := range slice {
		fmt.Print(items)
		p.sum = p.sum + items
	}

	fmt.Print("\n\n", p.sum, "\n")
}

func main() {

	var array []int

	for i := 0; i < 100; i++ {
		array = append(array, i)
	}
	fmt.Println(array)
	p := Point{0, 20, array, 0}
	k := Point{20, 40, array, 0}
	w := Point{40, 60, array, 0}
	q := Point{60, 80, array, 0}
	z := Point{80, 99, array, 0}

	go runtine1(p)
	runtine1(k)
	go runtine1(w)
	runtine1(q)
	runtine1(z)

}
