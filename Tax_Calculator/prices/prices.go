package prices

import (
	"fmt"

	"github.com/tax_calculator/iomanager"
)

type TaxIncludedPrices struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPrices) Process(done chan bool, errorVal chan error) {
	inputPrices, err := job.IOManager.ReadLines()

	if err != nil {
		errorVal <- err
	}

	job.TaxIncludedPrices = make(map[string]string, len(inputPrices))

	for _, price := range inputPrices {
		taxAddedPrice := fmt.Sprintf("%0.2f", price*(1+job.TaxRate))
		job.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = taxAddedPrice
	}

	job.IOManager.WriteJSON(job)

	done <- true
}

func NewTaxIncludedPrices(taxRate float64, iom iomanager.IOManager) *TaxIncludedPrices {
	return &TaxIncludedPrices{
		IOManager: iom,
		TaxRate:   taxRate,
	}
}
