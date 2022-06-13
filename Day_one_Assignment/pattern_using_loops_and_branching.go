
//*********
/*Print the following pattern using loops and branching:
1
2 6
3 7 10
4 8 11 13
5 9 12 14 15.
*/

package main

import "fmt"

func pattern(number int) {

	for i := 1; i <= number; i++ {
		var num = i
		for j := 1; j <= i; j++ {
			fmt.Print(num, " ")
			num = num + number - j
		}
		fmt.Println()
	}

}
func main() {
	var number int
	fmt.Println("Enter the Number :")
	fmt.Scanf("%d", &number)
	pattern(number)
}

