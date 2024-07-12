package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {

	title, body := getNoteData()

	userNote, err := note.New(title, body)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Print("Saving the note failed.")
		return
	}

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
