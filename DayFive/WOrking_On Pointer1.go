package main

import "fmt"

func change(str *string) {
	*str = " Rana Sanga"
}

func change1(str string) {
	str = " Babar"
}

func main() {
	str := " chatrapati sambaji"

	fmt.Println("King Who Kill Jangis Khan In Head To Head Fight",str)
	change(&str)

	fmt.Println("King Who Kill Jangis Khan In Head To Head Fight",str)
	
	change1(str)
	fmt.Println("King Who Kill Jangis Khan In Head To Head Fight",str)
}