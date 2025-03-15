package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	revenue = userEnteredValue("Revenue: ₹")
	expenses = userEnteredValue("Expenses: ₹")
	taxRate = userEnteredValue("Tax Rate: ")

	earningsBeforeTax, earningsAfterTax, ratio := calculations(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax: ₹%v\n", earningsBeforeTax)
	fmt.Printf("Earnings After Tax: ₹%v\n", earningsAfterTax)
	fmt.Printf("Ratio: %0.2f\n", ratio)
}

func userEnteredValue(text string) (value float64) {
	fmt.Print(text)
	fmt.Scan(&value)

	return value
}

func calculations(revenue, expenses, taxRate float64) (ebt, eat, ratio float64) {
	ebt = revenue - expenses
	eat = ebt * (1 - (taxRate / 100))
	ratio = ebt / eat

	return ebt, eat, ratio
}
