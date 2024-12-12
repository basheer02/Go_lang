package main

import "fmt"

func main() {
	/*age := 32

	var agePointer *int

	agePointer := &age

	fmt.Println(" Age :", *agePointer)
	getAdultYears(agePointer)
	fmt.Println("\n Adult since", age, "years ago")*/

	result := add(1, 3)

	res := add("a", "b")

	fmt.Println(result, res)
}

// to use any type of value, use type interface
func printSomething(value interface{}) {
	intVal, ok := value.(int) // it returns the value and it is true or false

	if ok {
		fmt.Println(" Integer :", intVal)
	}

	// switch value.(type) {
	// 	case int:
	// 		fmt.Println(" Integer :", value)
	// 	case float64:
	// 		fmt.Println(" Float :", value)
	// 	case string:
	// 		fmt.Println(" String :", value)
	// }
}

func add[T int|float64|string](a, b T) T { //to handle multiple type scenarios

	return a + b

}

/*func getAdultYears(age *int) {
	//return *age - 18

	*age -= 18
}*/
