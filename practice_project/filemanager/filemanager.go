package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Filemanager struct {
	inputPath  string
	outputPath string
}

func (fm Filemanager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.inputPath)

	if err != nil {
		return nil, errors.New(" could not open file ")
	}

	// defer keyword is used to be done something when the method completes its execution
	defer file.Close() // it close the file automatically when the method execution completes(whether the method contain error or not)

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		//file.Close()
		return nil, errors.New(" failed to read price ")
	}

	//file.Close()

	return lines, nil
}

func (fm Filemanager) WriteJSON(data interface{}) error {
	file, err := os.Create(fm.outputPath)

	if err != nil {
		return errors.New(" Failed to create file")
	}

	defer file.Close()

	time.Sleep(1 * time.Second) // for concurrency test

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		//file.Close()
		return errors.New(" Failed to convert data to JSON")
	}

	//file.Close()
	return nil
}

func New(input, output string) Filemanager {
	return Filemanager{
		inputPath:  input,
		outputPath: output,
	}
}
