package main

import (
	"fmt"
	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)



const accFileName = "balance.txt"

func main() {

	balance, err := fileops.GetDataFromFile(accFileName)
	//exit := true

	if err != nil {

		var cnt string
		fmt.Println(err)
		fmt.Print("\n Do you want to continue(y/n) : ")
		fmt.Scanln(&cnt)
		ch := cnt == "y"

		if !ch {
			//panic("Can't continue") #it also stop the execution
			return
		}
	}

	for /*exit*/ {

		options()

		//depositMoney := choice == 1
		//withDrawMoney := choice == 2
		//checkBalance := choice == 3
		//exit = choice != 0

		var amount int

		/*if depositMoney {
			fmt.Print("\n Deposit\n\n Enter deposit amount : ")
			fmt.Scan(&amount)
			balance += amount
			fmt.Println("\n Successfully deposited")
		} else if withDrawMoney {
			fmt.Print("\n Enter amount : ")
			fmt.Scan(&amount)

			if amount > balance {
				println("\n Insufficient balance")
				continue
			}

			balance -= amount
			fmt.Println("\n Amount Successfully debited")

		} else if checkBalance {
			fmt.Printf("\n Balance : %v\n", balance)
		} else {
			fmt.Println("\n Goodbye")
			break
		}*/

		switch choice {
		case 0:
			fmt.Println(" For help contact", randomdata.PhoneNumber())
			fmt.Println("\n Goodbye ")
			return
		case 1:
			fmt.Print("\n Deposit\n\n Enter deposit amount : ")
			fmt.Scan(&amount)
			balance += amount
			fileops.WriteDataToFile(balance, accFileName)
			fmt.Println("\n Successfully deposited")
		case 2:
			fmt.Print("\n Enter amount : ")
			fmt.Scan(&amount)

			if amount > balance {
				println("\n Insufficient balance")
				continue
			}

			balance -= amount
			fmt.Println("\n Amount Successfully debited")
			fileops.WriteDataToFile(balance, accFileName)
		case 3:
			fmt.Printf("\n Balance : %v\n", balance)
		default:
			fmt.Println("\n Wrong choice ")
		}
	}
}
