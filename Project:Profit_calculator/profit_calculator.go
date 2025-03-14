package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Revenue: ₹")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ₹")
	fmt.Scan(&expenses)

	fmt.Print("taxRate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (1 - (taxRate / 100))
	ratio := earningsBeforeTax / earningsAfterTax

	fmt.Println("Earnings Before Tax: ₹", earningsBeforeTax)
	fmt.Println("Earnings After Tax: ₹", earningsAfterTax)
	fmt.Println("Ratio: ", ratio)
}
