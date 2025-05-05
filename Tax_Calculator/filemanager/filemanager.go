package filemanager

import (
	"bufio"
	"encoding/json"
	"os"
	"strconv"
	"time"
)

type FileManager struct {
	InputPath  string
	OutputPath string
}

func (fm FileManager) ReadLines() ([]float64, error) {
	file, err := os.OpenFile(fm.InputPath, os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
	}

	defer file.Close() // Will be executed when the outer function completed execution

	scanner := bufio.NewScanner(file)

	list := []float64{}

	for scanner.Scan() {
		price, _ := strconv.ParseFloat(scanner.Text(), 64)

		list = append(list, price)
	}

	err = scanner.Err()

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (fm FileManager) WriteJSON(data interface{}) error {
	file, err := os.Create(fm.OutputPath)

	if err != nil {
		file.Close()
		return err
	}

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return err
	}

	file.Close()
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputPath:  inputPath,
		OutputPath: outputPath,
	}
}
