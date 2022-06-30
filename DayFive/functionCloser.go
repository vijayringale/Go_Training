package main

import "fmt"

func test(k int) {
	fmt.Println(" You Are in The Trubel :- " , k)
}


func main(){

	p := test
	p(23)


	tester := func( p  int ){

		fmt.Println("value in tester ",p)


	}
	tester(34)

}