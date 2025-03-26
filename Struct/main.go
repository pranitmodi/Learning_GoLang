package main

import (
	"fmt"
	"time"
)

// Struct - when mix of datatypes
// this is like defining a custom data type
type UserData struct {
	firstName string
	lastName  string
	email     string
	phone     int
}

var allUsers = make([]UserData, 0)

func main() {
	BookTicket("Alfred","Robert","ar@gmail.com",234567890)
	BookTicket("John", "Doe", "john.doe@example.com", 1234567890)

	// printing all users
	fmt.Println("Printing all users:")
	time.Sleep(3 * time.Second)
	// time.Sleep(1 * time.Second) is used to pause the execution of the program for 1 second, so that the output is printed after 1 second.

	for _, user := range allUsers {
		println("First Name:", user.firstName)
		println("Last Name:", user.lastName)
		println("Email:", user.email)
		println("Phone:", user.phone)
		println()
	}
}

func BookTicket(firstName string, lastName string, email string, phone int) {
	// creating a new user
	var newUser = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		phone:     phone,
	}

	// adding the new user to the list of all users
	allUsers = append(allUsers, newUser)
}