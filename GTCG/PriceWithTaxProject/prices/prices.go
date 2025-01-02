package prices

import (
	"fmt"
	"math"

	"github.com/tax/conversion"
	"github.com/tax/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData(fileName string) {
	lines, error := filemanager.ReadLines(fileName)

	if error != nil {
		fmt.Println("An error occurred while reading a file: ", error)
		return
	}

	prices, error := conversion.StringsToFloats(lines)

	if error != nil {
		fmt.Println("An error occurred while reading a file: ", error)
		return
	}

	job.InputPrices = prices

}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData("prices.txt")

	result := make(map[string]float64)

	for _, price := range job.InputPrices {

		fmt.Println("price ==> ", price)
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = math.Round(taxIncludedPrice*100) / 100
	}

	job.TaxIncludedPrices = result

	filemanager.WriteJSON("result.json", job)

	fmt.Println("result ==> ", result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}
