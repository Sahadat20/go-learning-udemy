package main

import (
	"fmt"
	"math"
)

func main() {

	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5
	fmt.Print("Entrt investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years:")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// fmt.Println("Future Value:", futureValue)

	formattedFV := fmt.Sprintf("Future Value: %v\n", futureValue)

	outputText(formattedFV)
	fmt.Printf("Future real value %.0f", futureRealValue)
	fmt.Print(`Multi Line 
		String`)
}

func outputText(text string) {
	fmt.Print(text)
}

// one waye to return value; (Recommended)
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fmt.Println("Calcualtion function")
	const inflationRate = 2.5
	fv := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	frv := fv / math.Pow(1+inflationRate/100, years)
	return fv, frv

}

// second waye to return value; 26. An Alternative Return Value Syntax
func calculateFutureValues2(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {
	fmt.Println("Calcualtion function")
	const inflationRate = 2.5
	fv = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	frv = fv / math.Pow(1+inflationRate/100, years)
	return

}
