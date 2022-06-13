
// Write a recursive function to compute the factorial
import "fmt"
func factorial(number int) int {
	if number > 1 {
		return number * factorial(number-1)
	} else{
		return 1
	}
}

func main() {
	var number int
	fmt.Println("Enter the Number :")
	fmt.Scanf("%d", &number)
	fmt.Println("Factorial of number", number, "is :", factorial(number))

}
