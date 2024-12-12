package prices

import (
	"fmt"

	"example.com/practice/conversion"
	"example.com/practice/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager        iomanager.IOManager `json:"-"`
	TaxRate          float64             `json:"tax_rate"`
	Prices           []float64           `json:"prices"`
	TaxIncludedPrice map[string]string   `json:"tax_included_price"`
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringToFloat(lines)

	if err != nil {
		return err
	}

	job.Prices = prices

	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChannel chan bool, errorChan chan error) {

	err := job.LoadData()

	//errorChan <- errors.New("Error")

	if err != nil {
		errorChan <- err
		return
	}

	result := make(map[string]string)

	for _, price := range job.Prices {
		taxIncludedprice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedprice)
	}

	job.TaxIncludedPrice = result
	fmt.Println(result)

	err = job.IOManager.WriteJSON(job)

	if err != nil {
		errorChan <- err
		return
	}
	doneChannel <- true
}

func New(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
		TaxRate:   taxRate,
	}
}
