package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser user
	// var appUser2 user
	// var appUser3 user
	var appUser4 *user.User
	// appUser = user{
	// 	firstname: firstName,
	// 	lastName:  lastName,
	// 	birthdate: birthdate,
	// 	createdAt: time.Now(),
	// }
	// appUser2 = user{ //shorter notation
	// 	firstName,
	// 	lastName,
	// 	birthdate,
	// 	time.Now(),
	// }
	// appUser3 = user{} //empty struct
	appUser4, err := user.NewUser(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println(err)
		return
	}
	admin := user.NewAdmin("test@gmail.com", "test1234")
	admin.OutputUserDetailsMethod()
	// ... do something awesome with that gathered data!

	// outputUserDetails(appUser)
	// outputUserDetailsPointer(&appUser2)
	// outputUserDetails(appUser3)
	// method call
	// appUser.outputUserDetailsMethod()
	// appUser.clearUserName()
	// appUser.outputUserDetailsMethod()
	// appUser.clearUserNamePointer()
	// appUser.outputUserDetailsMethod()
	appUser4.OutputUserDetailsMethod()
	appUser4.ClearUserNamePointer()
	appUser4.OutputUserDetailsMethod()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
