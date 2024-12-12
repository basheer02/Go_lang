package functions

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	dNumbers := transformNumbers(&numbers, getTransformerFunction(2))
	fmt.Println(dNumbers)
	tNumbers := transformNumbers(&numbers, getTransformerFunction(3))
	fmt.Println(tNumbers)
}

// passing function as a parameter
func transformNumbers(numbers *[]int, transform transformFn) []int { //instead of type alias can use transform func(int) int
	dNumbers := []int{}

	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}

	return dNumbers
}

// returning functions as values
func getTransformerFunction(number int) transformFn { //instead of type alias can use transform func(int) int
	if number == 2 {
		return double
	}
	return triple
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
