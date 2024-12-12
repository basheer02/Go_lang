package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title string `json:"title"` // metadata, when convert to json the field title as what we given
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New(" Invalid input ")
	}

	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ","_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note) // only publicly available content can encode to json


	if err != nil {
		return err
	} 

	return os.WriteFile(fileName, json, 0644)
	
}

func (note Note) Display() {
	fmt.Println(" Title:", note.Title)
	fmt.Println(" Content :", note.Content)
}