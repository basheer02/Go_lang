package main

import "fmt"


type ratingMap map[any]any

func (m ratingMap) Display(){
	fmt.Println(m)
}

func main() {
	usernames := make([]string, 2, 5) //type, reserved length, capacity(optional)

	usernames[0] = "qwer"

	usernames = append(usernames, "abcd", "bcda", "kjfjdf")
	usernames = append(usernames, "gdhjf") //if capacity is over, it realocates memory

	fmt.Println(usernames)

	ratings := make(ratingMap) //type, length(optional for map)

	ratings["go"] = 4.7
	ratings["react"] = 4.8
	ratings["node"] = 4.5

	//ratings.Display()

	for i, val := range usernames {
		fmt.Println(" Index :",i," Value :", val)
	}

	for key, val := range ratings {
		fmt.Println(key,"\t:", val)
	}
}