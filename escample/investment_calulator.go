package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

// integers and floats
func main() {
	var investAmount, expReturnRate, years float64

	fmt.Print(" Enter investment amount : ")
	fmt.Scan(&investAmount)

	fmt.Print(" Enter expected return rate : ")
	fmt.Scan(&expReturnRate)

	fmt.Print(" Enter no.of years : ")
	fmt.Scan(&years)

	fmt.Println("\n -------------------------")
	fmt.Println(" Investment Amount :", investAmount)
	fmt.Println(" Expected return Rate :", expReturnRate)
	fmt.Println(" No.of years :", years)

	amount, actualAmount := calcAmount(investAmount, expReturnRate, years)

	formattedAmount := fmt.Sprintf(" Amount : %.1f\n", amount)
	formattedActualAmount := fmt.Sprintf(" Amount as per inflation : %.1f\n", actualAmount)

	fmt.Print(formattedAmount, formattedActualAmount)
}

func calcAmount(investAmount, expReturnRate, years float64) /*(float64, float64)*/ (amount, actualAmount float64) {

	amount = investAmount * math.Pow(1+expReturnRate/100, years)
	actualAmount = amount / math.Pow(1+inflationRate/100, years)
	return
	//return amount, actualAmount
}

// strings

/*func main()  {
	var firstName, lastName string

	fmt.Print(" Enter first name : ")
	fmt.Scanln(&firstName)

	fmt.Print(" Enter last name : ")
	fmt.Scanln(&lastName)

	name := firstName + " " + lastName

	fmt.Println(" Name : ", name)
}

*/
