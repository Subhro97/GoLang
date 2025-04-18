package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/user_input/note"
	"github.com/user_input/todo"
)

// Defining Interface
type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

// Interface Embedding
type outputter interface {
	saver
	displayer
}

func main() {
	printValueUsingSwitch(1)
	printValueUsingSwitch(1.5)
	printValueUsingSwitch("Shree Ram!!")

	printAnyValue(2)
	printAnyValue(2.5)
	printAnyValue("Jai Shree Ram!!")

	intVal := genericFunction(2, 2)
	fmt.Println(intVal)

	floatVal := genericFunction(2.5, 2.5)
	fmt.Println(floatVal)

	stringVal := genericFunction("Jai ", "Bajrangbaali!!")
	fmt.Println(stringVal)

	title, content, err := getNoteData()

	if err != nil {
		fmt.Println(err)
	}

	note := note.New(title, content)

	err = outputData(note)

	if err != nil {
		fmt.Println(err)
		return
	}

	text, err := getUserInput("Enter Todo: ")

	if err != nil {
		fmt.Println(err)
	}

	todo := todo.New(text)
	outputData(todo)

}

func printValueUsingSwitch(value any) {
	switch value.(type) { // This type extraction only works for switch statements
	case int:
		fmt.Println("Integer Value: ", value)

	case float64:
		fmt.Println("Float Value: ", value)

	case string:
		fmt.Println("String Value: ", value)
	}
}

func printAnyValue(value any) {
	intValue, ok := value.(int)

	if ok {
		fmt.Println("Integer Value: ", intValue)
	}

	floatValue, ok := value.(float64)

	if ok {
		fmt.Println("Float Value: ", floatValue)
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("String Value: ", stringVal)
	}
}

func genericFunction[T string | int | float64](a, b T) T {
	return a + b
}

func outputData(data outputter) error {
	data.Display()
	dataErr := data.Save()

	if dataErr != nil {
		return dataErr
	}

	return nil
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
