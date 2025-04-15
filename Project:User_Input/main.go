package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/user_input/note"
)

func main() {
	title, content, err := getNoteData()

	if err != nil {
		fmt.Println(err)
	}

	note := note.New(title, content)
	note.Display()
	note.Save()
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note Title: ")

	if err != nil {
		return "", "", err
	}

	content, err := getUserInput("Note Content: ")

	if err != nil {
		return "", "", err
	}

	return title, content, err
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n') // Every line of Print when entered automatically ends with \n

	if err != nil {
		return "", err
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // Windows adds extra \r along with \n in line breaks

	return text, nil
}
