// Map - maps unique keys to values
package main

import "fmt"

func main() {
	var userData = make(map[string]string)
	userData["name"] = "John Doe"
	userData["email"] = "Jd@gmail.com"
	// userData["phone"] = 1234567890 - can't have phone number as ma can only hold same type of values, as phone number will be an integer, it will throw an error

	var allUsers = make([]map[string]string, 0)
	allUsers = append(allUsers, userData)

	fmt.Println(allUsers)
}