package main

import (
	// "fmt"

	"fmt"

	// "example.com/price-calculator/cmdmanager"
	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}
	channels := make([]chan bool, len(taxRates))
	erroChans := make([]chan error, len(taxRates))
	for index, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("results_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		// cmdm := cmdmanager.New()
		// priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		channels[index] = make(chan bool)
		erroChans[index] = make(chan error)
		go priceJob.Process(channels[index], erroChans[index])
		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	for index := range taxRates {
		select {
		case err := <-erroChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-channels[index]:
			fmt.Println("Done!")

		}
	}

	// for _, doneChan := range channels {
	// 	<-doneChan
	// }

}
