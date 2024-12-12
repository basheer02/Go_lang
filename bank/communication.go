package main

import (
	"fmt"
)

var choice int

func options() {

	fmt.Println("\n *****************")
	fmt.Println("\n Banking \n\n 1. Deposit\n 2. Withdraw\n 3. Check balance\n 0. Exit")
	fmt.Print("\n Enter your choice : ")
	fmt.Scan(&choice)
}
