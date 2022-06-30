package main

import "fmt"

func main() {

	toChange := "hello"
	var pointer *string = &toChange

	fmt.Println(pointer,"\t",&pointer,"\t",*pointer,"\t",toChange,"\t",&toChange,"\t",*&toChange)

}