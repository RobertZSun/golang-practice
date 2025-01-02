package main

import (
	"github.com/tax/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	// result := make(map[float64][]float64)

	for _, tax := range taxRates {
		priceJobPointer := prices.NewTaxIncludedPriceJob(tax)
		priceJobPointer.Process()
	}

}
