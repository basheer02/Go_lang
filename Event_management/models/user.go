package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"

	hashedPassword, err := utils.HashPassword(u.Password) // secure the password using hash

	if err != nil {
		return err
	}

	result, err := db.DB.Exec(query, u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, _ := result.LastInsertId()

	u.ID = userId

	return nil
}

func (u *User) ValidateCredentials() error {

	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&u.ID ,&retrievedPassword)

	if err != nil {
		return errors.New(" Invalid username")
	}

	checkPassword := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !checkPassword{
		return errors.New(" Invalid password")
	}

	return nil

}
