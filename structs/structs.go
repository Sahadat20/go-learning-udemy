package main

import (
	"fmt"
	"time"
)

type user struct {
	firstname string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user
	var appUser2 user
	var appUser3 user
	appUser = user{
		firstname: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
	appUser2 = user{ //shorter notation
		firstName,
		lastName,
		birthdate,
		time.Now(),
	}
	appUser3 = user{} //empty struct

	// ... do something awesome with that gathered data!

	outputUserDetails(appUser)
	outputUserDetailsPointer(&appUser2)
	outputUserDetails(appUser3)
}

func outputUserDetails(u user) {
	fmt.Println(u.firstname, u.lastName, u.birthdate, u.createdAt)
}

// just for practice
func outputUserDetailsPointer(u *user) {
	// u.firstname it is shorter version of (*u).firstname , that support by go
	fmt.Println(u.firstname, (*u).lastName, u.birthdate, u.createdAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
