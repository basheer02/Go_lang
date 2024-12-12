// tax calculator, input read from file, calculates it and write to file as json

package main

import (
	"fmt"

	"example.com/practice/filemanager"
	"example.com/practice/prices"
)

func main() {

	taxRates := []float64{0, .07, .1, .15}
	doneChannels := make([]chan bool, len(taxRates))
	errorChannels := make([]chan error, len(taxRates))

	for ind, tax := range taxRates {
		doneChannels[ind], errorChannels[ind] = make(chan bool), make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("tax_%.0f.json", tax*100))
		//cmd := cmdmanager.New()
		priceJob := prices.New(fm, tax)
		go priceJob.Process(doneChannels[ind], errorChannels[ind])

	}

	for index := range taxRates {
		select {
		case err := <-errorChannels[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChannels[index]:
			fmt.Println("done")
		}
	}

	/*for _, done := range doneChannels {
		<- done
	}*/

}
