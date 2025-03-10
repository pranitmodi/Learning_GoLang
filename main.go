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

	//if you don't want to use the index, just replace it with an underscore
	for index, task := range taskItems {
		//fmt.Println(index+1, ": ", task) //this adds extra spaces
		fmt.Printf("%d: %s\n", index+1, task)
	}

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
