package main

import "fmt"

func main() {
	age := 32 //regular variable

	var agePointer *int
	agePointer = &age

	fmt.Println("Age:", *agePointer)

	editAgeToAdultsYears(agePointer)
	// fmt.Println(adultYears)
	fmt.Println("Age:", age)
}

func editAgeToAdultsYears(age *int) {

	// return *age - 18
	*age = *age - 18
}
