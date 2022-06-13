
package main

import (
	"fmt"
	"math"
)

// *********
// Write and call a function to reverse a string using for loop.
// *********


func Reverse(s string) string {
	str1 := []byte(s)

	for i := 0; i < len(str1)/2; i++ {
		str1[len(str1)-i-1], str1[i] = str1[i], str1[len(str1)-i-1]
	}
	return string(str1)
}
func main() {
	fmt.Println(Reverse("tnairox"))
}
