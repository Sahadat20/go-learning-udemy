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

type Admin struct {
	email    string
	password string
	User     //embedded the User struct here
}

// method
func (u User) OutputUserDetailsMethod() {
	fmt.Println(u.firstname, u.lastName, u.birthdate, u.createdAt)
}
func (u User) ClearUserName() { // not change the username
	u.firstname = ""
	u.lastName = ""
}
func (u *User) ClearUserNamePointer() { // change the username
	u.firstname = ""
	u.lastName = ""
}

// a constructor that build struct
func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name , last name and birthday are required")
	}
	return &User{
		firstname: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstname: "Admin",
			lastName:  "User",
			birthdate: "-----",
			createdAt: time.Now(),
		},
	}
}

func outputUserDetails(u User) {
	fmt.Println(u.firstname, u.lastName, u.birthdate, u.createdAt)
}

// just for practice
func OutputUserDetailsPointer(u *User) {
	// u.firstname it is shorter version of (*u).firstname , that support by go
	fmt.Println(u.firstname, (*u).lastName, u.birthdate, u.createdAt)
}
