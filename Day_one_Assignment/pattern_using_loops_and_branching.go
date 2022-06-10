
package main

import (
	"fmt"
	"math"
)
//*********
/*Print the following pattern using loops and branching:
1
2 6
3 7 10
4 8 11 13
5 9 12 14 15.
*/

func main() {
	var num = 5
	lowerLeftCorner := num*(num-1)/2 + 1
	lastInColumn := lowerLeftCorner
	lastInRow := 1
	for i, row := 1, 1; row <= num; i++ {
		w := len(fmt.Sprint(lastInColumn))
		if i < lastInRow {
			fmt.Printf("%*d ", w, i)
			lastInColumn++
		} else {
			fmt.Printf("%*d\n", w, i)
			row++
			lastInRow += row
			lastInColumn = lowerLeftCorner
		}
	}
}
