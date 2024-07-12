package note

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/pterm/pterm"
)

type Note struct {
	Title     string
	Body      string
	CreatedAt time.Time
}

func (n Note) Display() {
	pterm.Printf(pterm.Black("\n-------------------------\n\n"))
	pterm.Println(pterm.Black("Title: ") + pterm.Yellow(n.Title) + ("\n"))
	pterm.Println(pterm.Black("Content:") + ("\n"))
	pterm.Printfln(n.Body)
	pterm.Printf(pterm.Black("\n-------------------------"))
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

// Method [Create New Note]
func New(title, body string) (Note, error) {
	if title == "" || body == "" {
		return Note{}, errors.New("title and body inputs are required")
	}
	return Note{
		Title:     title,
		Body:      body,
		CreatedAt: time.Now(),
	}, nil
}
