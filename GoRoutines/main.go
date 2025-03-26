package main

// By default, sequential code execution is done in GoLang.
// But, if you want to run multiple tasks concurrently, you can use GoRoutines.
// GoRoutines are lightweight threads managed by the Go runtime.
// GoRoutines are used to run multiple tasks concurrently.
// Can be created by using the keyword go followed by the function name.

// Concurreny in Go is cheap and easy

// *If the thread where you have used go routine is not finished, but the rest of the lines are executed,
// then the program will exit without waiting for the go routine to finish.

// WaitGroup
// WaitGroup is used to wait for the execution of a collection of goroutines.
// Add - Add the number of goroutines to wait for.
// Done - Decrements the count of the waitgroup.
// Wait - Blocks until the WaitGroup counter is zero.

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(1)
	go printNumbers() // For rhis function to execute completely and stopping program from execting, WaitGroup is used.
	fmt.Println("Hello, code should be executing concurrently")

	wg.Wait()
}

func printNumbers() {
	time.Sleep(5 * time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	wg.Done()
}