package main

import "fmt"

// first class function main()
func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := doubleNumbers(&numbers)
	doubled_n := transformNumbers(&numbers, getTrasformerFunction(2))
	triple_n := transformNumbers(&numbers, getTrasformerFunction(3))
	triple_n1 := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(doubled_n)
	fmt.Println(triple_n)
	fmt.Println(triple_n1)
	// fmt.Println(numbers)

	// anonymous function

}

func transformNumbers(number *[]int, transform func(int) int) []int {
	dnumbers := []int{}
	firstVal := (*number)[0]
	fmt.Println(firstVal)
	(*number)[0] = 5
	for _, val := range *number {

		dnumbers = append(dnumbers, transform(val))

	}

	return dnumbers
}

func getTrasformerFunction(jumVal int) func(int) int {
	if jumVal == 2 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}
func triple(number int) int {
	return number * 3
}
func doubleNumbers(number *[]int) []int {
	dnumbers := []int{}
	for _, val := range *number {

		dnumbers = append(dnumbers, val*2)

	}

	return dnumbers
}
