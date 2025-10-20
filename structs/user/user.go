package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
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

func outputUserDetails(u user) {
	fmt.Println(u.firstname, u.lastName, u.birthdate, u.createdAt)
}

// just for practice
func outputUserDetailsPointer(u *user) {
	// u.firstname it is shorter version of (*u).firstname , that support by go
	fmt.Println(u.firstname, (*u).lastName, u.birthdate, u.createdAt)
}
