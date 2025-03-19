package main

import "fmt"

func main() {
	var shortGolang = "Watch Go crash course"
	var fullGolang = "Watch Nana's Golang Full Course"
	desertName := "donut"
	rewardDessert := "Reward myself with a " + desertName

	taskItems := []string{shortGolang, fullGolang, rewardDessert}
	// Adding size to this slice, make it an array

	fmt.Println("###### Welcome to our Todolist App! ######")

	fmt.Println("List of my Todos")
	//fmt.Println(taskItems)
	printTasks(taskItems)

	// Adding a new Task
	fmt.Println()
	fmt.Println("Tasks after adding a new task:")
	taskItems = addTask(taskItems, "Add a new task")
	printTasks(taskItems)

	fmt.Println()
	fmt.Println("Tutorials")
	fmt.Println(shortGolang)
	fmt.Println(fullGolang)

	fmt.Println()
	fmt.Println("Rewards")
	fmt.Println(rewardDessert)

	fmt.Println()
	fmt.Println("My Project")
	fmt.Println(fullGolang)
}

func printTasks(taskItems []string) {
	//if you don't want to use the index, just replace it with an underscore
	for index, task := range taskItems {
		//fmt.Println(index+1, ": ", task) //this adds extra spaces
		fmt.Printf("%d: %s\n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	return append(taskItems, newTask)
}
