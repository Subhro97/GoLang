package main

import (
	"fmt"

	"github.com/tax_calculator/cmdmanager"
	"github.com/tax_calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("taxed_price_%.0f.json", taxRate*100))
		cmd := cmdmanager.New()

		taxIncludedPriceJob := prices.NewTaxIncludedPrices(taxRate, cmd)
		err := taxIncludedPriceJob.Process()

		if err != nil {
			fmt.Println("Unable to Process Job!")
			return
		}
	}

}

/*
1. We will read the prices from a file
2. based on the prices we will calculate the taxes include prices
3. Write these prices in the  JSON file
*/
