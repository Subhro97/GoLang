package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"content"`
}

func (todo Todo) Display() {
	fmt.Println("Your todo: ", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(content string) Todo {
	return Todo{
		Text: content,
	}
}
