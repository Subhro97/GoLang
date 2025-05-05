package main

import (
	"fmt"

	"github.com/tax_calculator/filemanager"
	"github.com/tax_calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index := range doneChans {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
	}

	for index, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("taxed_price_%.0f.json", taxRate*100))
		// cmd := cmdmanager.New()

		taxIncludedPriceJob := prices.NewTaxIncludedPrices(taxRate, fm)
		go taxIncludedPriceJob.Process(doneChans[index], errorChans[index])

	}

	for index := range taxRates {
		select { // Any one of the channels resolves will complete the execution of that routine
		case err := <-errorChans[index]:
			fmt.Println("Error Occured!!", err)

		case <-doneChans[index]:
			fmt.Println("Done!!")
		}
	}

}

/*
1. We will read the prices from a file
2. based on the prices we will calculate the taxes include prices
3. Write these prices in the  JSON file
*/
