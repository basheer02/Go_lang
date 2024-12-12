package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var key = "secretkey"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString([]byte(key))
}

func VerifyToken(token string) (int64, error) {

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) { //Parse is used to parse and validate token
		//This is a callback function that provides the key used to validate the JWT's signature. 
		//The jwt.Parse function will call this function during the parsing process.
		_, ok := t.Method.(*jwt.SigningMethodHMAC) // checking the signing method type is correct or not

		if !ok {
			return nil, errors.New(" Unexpected signing method")
		}

		return []byte(key), nil
	})

	if err != nil {
		return 0, errors.New(" Could not parse token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New(" Invalid token ")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New(" Invalid token claims ")
	}

	//email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil
}
