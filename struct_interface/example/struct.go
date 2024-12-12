package main

import (
	"fmt"
	//"example.com/struct/user"
)

type str string // custom type, same as struct(also srtuct operations)

func (text str) log() { 
	fmt.Println(text)
}

func main() {

	var name str = "abcd"
	name.log()

	/*userFirstName := getInput(" Enter first name : ")
	userLastName := getInput(" Enter lat name : ")
	userDOB := getInput(" Enter date of birth (MM/DD/YYYY) : ")

	var userData *user.User
	userData, err := user.New(userFirstName, userLastName, userDOB)

	admin := user.NewAdmin("abc@c.som", "123")

	admin.PrintUserData()

	if err != nil {
		fmt.Print(err)
		return
	}

	userData.PrintUserData()
	userData.ClearData()
	fmt.Println(" data cleared")
	userData.PrintUserData()*/
}

/*func printUserData(u *user){
	fmt.Println(" Name : ",u.firstName + " " + u.lastName) // (*u).firstName
	fmt.Println(" dob : ",u.dateOfBirth)
	fmt.Println(" created at : ",u.createdAt)
}*/

/*func getInput(input string) string {
	fmt.Print(input)
	var value string
	fmt.Scanln(&value)
	return value
}*/
