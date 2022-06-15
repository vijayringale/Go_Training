package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)

}

func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	fmt.Println("Poped Element", toRemove)
	return l
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(500)
	myStack.Push(300)
	myStack.Push(500)
	myStack.Push(700)
	myStack.Push(800)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)

}
