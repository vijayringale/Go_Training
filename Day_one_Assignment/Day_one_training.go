package main

// import (
// 	"fmt"
// 	"math"
// )

//*********
//Write and call a function to reverse a string using for loop.
//*********

// func Reverse(s string) (result string) {

// 	for i:=len(s)-1; i>=0; i--{
// 		result =result+string(s[i])
// 	}
// 	fmt.Println(result)

// 	return
//   }

// func main(){
// 	Reverse("tnairox")
// }

//*********
/*Print the following pattern using loops and branching:
1
2 6
3 7 10
4 8 11 13
5 9 12 14 15.
*/

// func main() {
// 	var num = 5
// 	lowerLeftCorner := num*(num-1)/2 + 1
// 	lastInColumn := lowerLeftCorner
// 	lastInRow := 1
// 	for i, row := 1, 1; row <= num; i++ {
// 		w := len(fmt.Sprint(lastInColumn))
// 		if i < lastInRow {
// 			fmt.Printf("%*d ", w, i)
// 			lastInColumn++
// 		} else {
// 			fmt.Printf("%*d\n", w, i)
// 			row++
// 			lastInRow += row
// 			lastInColumn = lowerLeftCorner
// 		}
// 	}
// }

// *******
// sorting/searching algorithm of your choice using Go.
//*******

// import "fmt"

// func margSort(arr []int) []int {
// 	if len(arr) < 2 {
// 		return arr
// 	}
// 	// Arr is array which contain the sorting values
// 	var middle int

// 	middle = (len(arr)) / 2

// 	return JoinArray(margSort(arr[:middle]), margSort(arr[middle:]))

// }

// func JoinArray(leftArr []int, rightArr []int) []int {
// 	var num int
// 	var i int
// 	var j int
// 	num, i, j = len(leftArr)+len(rightArr), 0, 0
// 	var array []int
// 	array = make([]int, num, num)

// 	var k int
// 	for k = 0; k < num; k++ {
// 		if i > len(leftArr)-1 && j <= len(rightArr)-1 {
// 			array[k] = rightArr[j]
// 			j++
// 		} else if j > len(rightArr)-1 && i <= len(leftArr)-1 {
// 			array[k] = leftArr[i]
// 			i++
// 		} else if leftArr[i] < rightArr[j] {
// 			array[k] = leftArr[i]
// 			i++
// 		} else {
// 			array[k] = rightArr[j]
// 			j++
// 		}
// 	}
// 	return array

// }

// func main() {
// 	var elements []int
// 	elements = []int{11, 4, 18, 6, 19, 21, 71, 13, 15, 2}
// 	fmt.Println("\n Before SOrting \n\n", elements)
// 	fmt.Println("\n After Sorting \n\n", margSort(elements))
// }

// *******
// Write a recursive function to compute the factorial(Optional).
//******88

import "fmt"

var factVal uint64 = 1

var i int = 1
var n int

func factorial(n int) uint64 {
	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= uint64(i)
		}

	}
	return factVal
}

func main() {
	fmt.Print("Enter a positive integer between 0 - 50 : ")
	fmt.Scan(&n)
	fmt.Print("Factorial is: ", factorial(n))
}
