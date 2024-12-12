package anonymousfunctions

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	dNumbers := transformNumbers(&numbers, func(number int) int{ // it's called anonymous function
		return number * 2
	})
	fmt.Println(dNumbers)
	tNumbers := transformNumbers(&numbers, func(number int) int{
		return number * 3
	})
	fmt.Println(tNumbers)

	double := createTransformer(2)
	triple := createTransformer(3)

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(doubled, tripled)
}

// passing function as a parameter
func transformNumbers(numbers *[]int, transform transformFn) []int { //instead of type alias can use transform func(int) int
	dNumbers := []int{}

	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}

	return dNumbers
}

func createTransformer(factor int) transformFn {
	return func(number int) int {
		return number * factor
	}
}
