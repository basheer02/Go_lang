package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface { //if interface only contain a single method name the interface as adding 'er' to the method name
	Save() error
}

type outputtable interface {
	Display()
	saver
}

func main() {

	/*file, _ := os.Open("todo.json")

	decoder := json.NewDecoder(file)

	var result map[string]interface{} // string/json.Number...represents json number as a string

	err := decoder.Decode(&result)

	if err != nil {
		fmt.Println(err)
		return
	}

	for key, val := range result {
		fmt.Printf(" %v : %v\n", key, val)
	}

	fmt.Println(result)*/
	todoUser, err := todo.New(getUserInput(" Enter todo text : "))

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(getNoteData())

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todoUser)
	if err != nil {
		return
	}

	outputData(userNote)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {

	err := data.Save()

	if err != nil {
		fmt.Println(" Saving the note failed.", err)
		return err
	}

	fmt.Println(" Note successfully saved")
	return nil
}

func getNoteData() (string, string) {

	title := getUserInput(" Enter title of the note : ")
	content := getUserInput(" Enter content : ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin) // to handle long text

	text, err := reader.ReadString('\n') // stop reading when \n occurs(new line)

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // sometimes at the end is \r\n

	return text
}
