package main

import (
	"fmt"
	"strings"
)

func main() {
	// printAll()
	// startServer()
	forLoop()	
}

func forLoop() {
	var i = 10
	for i > 0 { // syntax of for loop kinda like while loop in other languages
		fmt.Println(i)
		i--
	}
}

func printAll() {
	var shortGolang string = "Watch Go crash course"
	var fullGolang string = "Watch Nana's Golang Full Course"
	desertName := "donut"
	rewardDessert := "Reward myself with a " + desertName

	// Declaring a variable
	var number1 int
	fmt.Println("Default value of int:",number1) // prints 0 because it's the default value
	fmt.Printf("Type of number1: %T\n", number1) // %T is used to print the type of the variable 

	taskItems := []string{shortGolang, fullGolang, rewardDessert}
	// Adding size to this slice, make it an array

	fmt.Println("###### Welcome to our Todolist App! ######")

	//fmt.Println("List of my Todos")
	//fmt.Println(taskItems)
	fmt.Printf("List of my Todos: %v\n", taskItems)

	// Adding a new Task
	taskItems = addTask(taskItems)
	fmt.Println("Tasks after adding a new task:")
	printTasks(taskItems)

	var myFullName string = "Nana Kwame"
	for _, word := range strings.Fields(myFullName) { //this splits the string into words
		fmt.Println(word)
	}
}

func printTasks(taskItems []string) {
	//if you don't want to use the index, just replace it with an underscore
	for index, task := range taskItems {
		//fmt.Println(index+1, ": ", task) //this adds extra spaces
		fmt.Printf("%d: %s\n", index+1, task)
	}
}

func addTask(taskItems []string) []string {
	fmt.Print("Enter the new task: ")
	var newTask string
	fmt.Scanln(&newTask)
	return append(taskItems, newTask) //append returns a new slice
}
