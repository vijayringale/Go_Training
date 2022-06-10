
package main

import (
	"fmt"
	"math"
)

*********
Write and call a function to reverse a string using for loop.
*********

func Reverse(s string) (result string) {

	for i:=len(s)-1; i>=0; i--{
		result =result+string(s[i])
	}
	fmt.Println(result)

	return
  }

func main(){
	Reverse("tnairox")
}
