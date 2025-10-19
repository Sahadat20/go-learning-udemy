package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

// 28. Onwards to Control Structures
func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("Error")
		fmt.Print(err)

		fmt.Println("------")
		panic("Can't continue, sorry")
	}
	fmt.Println("Welcome to Go Bank!")
	fmt.Println(randomdata.PhoneNumber())
	// for i := 0; i < 200; i++ {
	for { //for infinite loop
		presentOptions()
		var choice int
		fmt.Print("Enter your choice:")
		fmt.Scan(&choice)

		fmt.Println("your choice is:", choice)

		if choice == 1 {
			fmt.Println("Your balance is ", accountBalance)
		} else if choice == 2 {
			var depositAmount float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid! amount must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance updated. Your new balance is ", accountBalance)
		} else if choice == 3 {
			var withdrawAmount float64
			fmt.Print("Your withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount! you can't withdraw more than you owe")
				continue
			}
			accountBalance -= withdrawAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance updated. Your new balance is ", accountBalance)
		} else {
			fmt.Println("Goodbye")
			break
		}
	}
	fmt.Println("Thanks for choosing our bank")

}
