package cmdmanager

import (
	"fmt"
	"strconv"
)

type CMDManager struct {
}

func (cmd CMDManager) ReadLines() ([]float64, error) {
	list := []float64{}

	for {
		var value string

		fmt.Print("Please enter price: ")
		fmt.Scanln(&value)

		price, _ := strconv.ParseFloat(value, 64)

		list = append(list, price)

		if value == "0" {
			break
		}
	}

	return list, nil
}

func (cmd CMDManager) WriteJSON(data interface{}) error {
	fmt.Print(data)

	return nil
}

func New() CMDManager {
	return CMDManager{}
}
