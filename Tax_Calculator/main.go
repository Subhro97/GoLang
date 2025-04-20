package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	prices := []float64{}
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	taxesIncludedPrices := map[string][]string{}

	file, err := os.OpenFile("prices.txt", os.O_RDONLY, 0644)

	if err != nil {
		fmt.Println("Failed to Read File!")
		file.Close()
		return
	}

	scanner := bufio.NewScanner(file)

	err = scanner.Err()

	if err != nil {
		fmt.Println("Failed to Scan the File!")
		return
	}

	for scanner.Scan() {
		price, _ := strconv.ParseFloat(scanner.Text(), 64)

		prices = append(prices, price)
	}

	for _, taxRate := range taxRates {
		modifiedPricesWithTax := make([]string, len(prices))

		for priceIndex, price := range prices {
			taxAddedPrice := fmt.Sprintf("%0.2f", price*(1+taxRate))
			modifiedPricesWithTax[priceIndex] = taxAddedPrice
		}

		taxesIncludedPrices[fmt.Sprintf("%0.2f", taxRate)] = modifiedPricesWithTax
	}

	fmt.Println(taxesIncludedPrices)

}

/*
1. We will read the prices from a file
2. based on the prices we will calculate the taxes include prices
3. Write these prices in the  JSON file
*/
