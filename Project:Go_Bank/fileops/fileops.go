package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadFloatFromFile(fileName string) (float64, error) {
	value, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Unable to read from file.")
		return 1000, err
	}

	valueText := string(value)

	readValue, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		fmt.Println("Unable to convert value to float.")

		return 1000, err
	}

	return readValue, nil

}

func WriteFloatToFile(fileName string, value float64) error {
	valueText := strconv.FormatFloat(value, 'f', -1, 64)

	err := os.WriteFile(fileName, []byte(valueText), 0644) // File Permission Code - Linux Specific

	if err != nil {
		return errors.New("failed to write to file")
	}

	return nil
}
