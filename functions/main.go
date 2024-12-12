package main

import "fmt"

func main() {

	numbers := []int{1, 19, 15, 16}
	sum := sumup(1,2,3,4,5,6,7)
	numbersSum := sumup(numbers...)

	fmt.Println(sum, numbersSum)
}

// variadic function
func sumup(numbers ...int) int { // can accept many paramaters
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
