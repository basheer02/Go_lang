package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (Todo, error) {

	if content == "" {
		return Todo{}, errors.New(" Invalid input ")
	}

	return Todo{
		Text: content,
	}, nil
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo) // only publicly available content can encode to json

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)

}

func (todo Todo) Display() {
	fmt.Println(" Todo :",todo.Text)
}
