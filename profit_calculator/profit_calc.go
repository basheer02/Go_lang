package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue,err1 := getInput(" Enter revenue : ")
	expenses,err2 := getInput(" Enter expenses : ") 
	taxRate,err3 := getInput(" Enter tax rate : ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

	ebt, profit, ratio := calcProfit(revenue, expenses, taxRate)
	fmt.Println(" Earnings before tax : ", ebt)
	fmt.Println(" Earnings after tax (Profit) : ", profit)
	fmt.Println(" ratio : ", ratio)

	storeResults(ebt, profit, ratio)
}

func getInput(text string) (float64, error) {

	var value float64
	fmt.Print(text)
	fmt.Scan(&value)

	if value <= 0 {
		return 0, errors.New(" Value Must be positive")
	}

	return value, nil
}

func calcProfit(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT : %.1f\nProfit : %.1f\nRatio : %.1f", ebt, profit, ratio)
	os.WriteFile("revenue.txt", []byte(results), 0644)
}
