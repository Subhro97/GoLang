// Functions, Strings and Formats

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, err := userEnteredValue("Revenue: ₹")

	if err != nil {
		panic(err)
	}

	expenses, err := userEnteredValue("Expenses: ₹")

	if err != nil {
		panic(err)
	}

	taxRate, err := userEnteredValue("Tax Rate: ")

	if err != nil {
		panic(err)
	}

	earningsBeforeTax, earningsAfterTax, ratio := calculations(revenue, expenses, taxRate)

	calErr := writeCalculationToFile(earningsBeforeTax, earningsAfterTax, ratio)

	if calErr != nil {
		panic(calErr)
	}

	fmt.Printf("Earnings Before Tax: ₹%v\n", earningsBeforeTax)
	fmt.Printf("Earnings After Tax: ₹%v\n", earningsAfterTax)
	fmt.Printf("Ratio: %0.2f\n", ratio)
}

func userEnteredValue(text string) (value float64, err error) {
	fmt.Print(text)
	fmt.Scan(&value)

	if value <= 0 {
		fmt.Println("Please Enter value greater than 0")
		return 0, errors.New("value is negative")
	}

	return value, nil
}

func calculations(revenue, expenses, taxRate float64) (ebt, eat, ratio float64) {
	ebt = revenue - expenses
	eat = ebt * (1 - (taxRate / 100))
	ratio = ebt / eat

	return ebt, eat, ratio
}

func writeCalculationToFile(ebt, eat, ratio float64) error {
	calculationText := fmt.Sprintf("EBT: ₹%0.2f\nEAT: ₹%0.2f\nRatio: ₹%0.2f", ebt, eat, ratio)

	err := os.WriteFile("Calculations.txt", []byte(calculationText), 0644)

	if err != nil {
		fmt.Println("Failed to write to File.")
		return err
	}
	return nil
}
