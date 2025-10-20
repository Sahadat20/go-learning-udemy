package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstname string
	lastName  string
	birthdate string
	createdAt time.Time
}

// method
func (u user) outputUserDetailsMethod() {
	fmt.Println(u.firstname, u.lastName, u.birthdate, u.createdAt)
}
func (u user) clearUserName() { // not change the username
	u.firstname = ""
	u.lastName = ""
}
func (u *user) clearUserNamePointer() { // change the username
	u.firstname = ""
	u.lastName = ""
}

// a constructor that build struct
func newUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name , last name and birthday are required")
	}
	return &user{
		firstname: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user
	var appUser2 user
	var appUser3 user
	var appUser4 *user
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
	appUser4, err := newUser(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println(err)
	}
	// ... do something awesome with that gathered data!

	outputUserDetails(appUser)
	outputUserDetailsPointer(&appUser2)
	outputUserDetails(appUser3)
	// method call
	appUser.outputUserDetailsMethod()
	appUser.clearUserName()
	appUser.outputUserDetailsMethod()
	appUser.clearUserNamePointer()
	appUser.outputUserDetailsMethod()
	appUser4.outputUserDetailsMethod()
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
	fmt.Scanln(&value)
	return value
}
