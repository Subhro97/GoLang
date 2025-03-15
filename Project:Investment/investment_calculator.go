package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount float64 = 1000.0
	years := 10.0
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scanln(&investmentAmount)
	fmt.Print("Years: ")
	fmt.Scanln(&years)
	fmt.Print("Expected Return Rate: ")
	fmt.Scanln(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
