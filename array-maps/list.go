package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) Output() {
	fmt.Println(m)
}

func main() {
	var productnames [4]string = [4]string{"A book"}
	prices := [4]float64{10.00, 0.33, 324.33, 20.0}
	fmt.Println(prices)
	fmt.Println(productnames)

	fmt.Println(prices[2])
	featuredPrices := prices[1:]
	highlitedPrice := featuredPrices[:1]
	fmt.Println(len(highlitedPrice), cap(highlitedPrice))
	fmt.Println(highlitedPrice[:1])
	fmt.Println(highlitedPrice[:3])

	// dynamic array

	pr := []float64{10.99, 8.99}
	pr = append(pr, 45.55, 10, 5)
	fmt.Println(pr)

	discountPrices := []float64{10.2, 52.5}

	// marge prices and discountPricescons
	pr = append(pr, discountPrices...)
	fmt.Println(pr)

	// given also dynamic but it ensure some memory management
	userNames := make([]string, 2, 5) //size, capacity
	userNames = append(userNames, "hello")
	userNames = append(userNames, "world")
	fmt.Println(userNames)

	// for map
	courseRatings := make(map[string]float64, 3)
	fmt.Println(courseRatings)
	newRatings := make(floatMap, 3)
	newRatings["panda"] = 474.3
	newRatings.Output()

	// for loop
	for index, value := range userNames {
		fmt.Println(index)
		fmt.Println(value)
	}

}
