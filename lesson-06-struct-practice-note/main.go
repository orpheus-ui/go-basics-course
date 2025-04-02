package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// "er" common convension for interface
type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

// type outputtable interface{
// 	Save() error
// 	Display()
// }

func main() {

	title, body := getNoteData()
	todoText := getUserInput("Todo text:")

	userNote, err := note.New(title, body)
	if err != nil {
		fmt.Println(err)
		return
	}

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Note Display, Save
	err = outputData(userNote)
	if err != nil {
		return
	}

	// Todo Display, Save
	err = outputData(todo)
	if err != nil {
		return
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Print("Saving the note failed.")
		return err
	}
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Title:")
	body := getUserInput("Note Body:")
	return title, body
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	// fmt.Scanln don't work for Longer text input
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	// Because in windows, a line break is \r\n
	text = strings.TrimSuffix(text, "\r")

	return text
}
