package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	dateOfBirth string
	createdAt   time.Time
}

type Admin struct {
	username string
	password string
	User
}


func NewAdmin(username, password string) Admin {
	return Admin{
		username: username,
		password: password,
		User: User{
			firstName: "admin",
			lastName: "admin",
			dateOfBirth: "---",
			createdAt: time.Now(),
		},
	}
}


func (u *User) PrintUserData() { // this method is now a field of struct, use pointer to save memory
	fmt.Println(" Name :", u.firstName+" "+u.lastName) // (*u).firstName
	fmt.Println(" dob :", u.dateOfBirth)
}

func (u *User) ClearData() {
	u.firstName = ""
	u.lastName = ""
	u.dateOfBirth = ""
}

func New(firstName, lastName, dob string) (*User, error) {

	if firstName == "" || lastName == "" || dob == "" {
		return nil, errors.New(" input should not be empty ")
	}

	return &User{
		firstName:   firstName,
		lastName:    lastName,
		dateOfBirth: dob,
		createdAt:   time.Now(),
	}, nil
}
