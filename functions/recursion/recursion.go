package recursion

import (
	"fmt"
)

func main() {
	fmt.Print(" Enter a number to find factorial : ")
	var num int
	fmt.Scanln(&num)
	fact := factorial(num)
	fmt.Printf(" factorial of %v : %v\n ", num, fact)
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}

	return number * factorial(number-1)
}
