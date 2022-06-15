package main

import "fmt"

type Point struct {
	x int
	y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func modifyValue(p Point) {
	p.x = 10
}

func modifyPointer(p *Point) {
	p.x = 5
	p.y = 5
}

func modifyReference(p *Point) {
	p = &Point{5, 5}
}

func main() {
	p := Point{0, 0}
	fmt.Println(p)

	modifyValue(p)
	fmt.Println(p)

	modifyPointer(&p)
	fmt.Println(p)

	p = Point{0, 0}
	modifyReference(&p)
	fmt.Println(p)
}
