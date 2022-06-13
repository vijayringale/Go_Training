mport "fmt"

func margSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	// Arr is array which contain the sorting values
	var middle int

	middle = (len(arr)) / 2

	return JoinArray(margSort(arr[:middle]), margSort(arr[middle:]))

}

func JoinArray(leftArr []int, rightArr []int) []int {
	var num int
	var i int
	var j int
	num, i, j = len(leftArr)+len(rightArr), 0, 0
	var array []int
	array = make([]int, num, num)

	var k int
	for k = 0; k < num; k++ {
		if i > len(leftArr)-1 && j <= len(rightArr)-1 {
			array[k] = rightArr[j]
			j++
		} else if j > len(rightArr)-1 && i <= len(leftArr)-1 {
			array[k] = leftArr[i]
			i++
		} else if leftArr[i] < rightArr[j] {
			array[k] = leftArr[i]
			i++
		} else {
			array[k] = rightArr[j]
			j++
		}
	}
	return array

}

func main() {
	var elements []int
	elements = []int{11, 4, 18, 6, 19, 21, 71, 13, 15, 2}
	fmt.Println("\n Before SOrting \n\n", elements)
	fmt.Println("\n After Sorting \n\n", margSort(elements))
}
