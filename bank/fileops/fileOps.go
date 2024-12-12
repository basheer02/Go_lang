package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetDataFromFile(fileName string) (int, error) {

	valueStr, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New(" Failed to fetch file")
	}

	value, err := strconv.Atoi(string(valueStr))
	if err != nil {
		return 0, errors.New(" Failed to fetch data")
	}

	return value, nil
}

func WriteDataToFile(value int, fileName string) {

	valueStr := fmt.Sprint(value)

	os.WriteFile(fileName, []byte(valueStr), 0644)

}
