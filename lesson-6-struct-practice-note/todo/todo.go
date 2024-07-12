package todo

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/pterm/pterm"
)

type Todo struct {
	Text string `json:"text"`
}

func (t Todo) Display() {
	pterm.Println(pterm.Black("\n\nTodo: ") + pterm.Yellow(t.Text))
	pterm.Println(pterm.Black("\n-------------------------"))
}

func (t Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(t)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

// Method [Create New Note]
func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("title and body inputs are required")
	}
	return Todo{
		Text: content,
	}, nil
}
