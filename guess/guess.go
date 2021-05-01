package main

import (
	"bufio"
	"fmt"
	"os"
)

// A number guessing program
func main() {
	// butfio reads things from input sources
	// os.Stdin means input source from console
	scanner := bufio.NewScanner((os.Stdin))

	low := 1
	high := 100

	fmt.Printf("Please think of a number between %v and %v", low, high)
	fmt.Println("\nPrint ENTER when ready")
	scanner.Scan() // wait for input

	// guess := 50
	for {
		guess := (low + high) / 2
		fmt.Println("My guess:", guess)
		fmt.Println("Is that:")
		fmt.Println("(a) Too high")
		fmt.Println("(b) Too low")
		fmt.Println("(c) Correct")

		scanner.Scan()
		response := scanner.Text() // returns what the user types

		// Binary Search Strategy

		// if response == "a" {
		// 	guess--
		// } else if response == "b" {
		// 	guess++
		// } else if response == "c" {
		// 	fmt.Println("I won!")
		// 	break // exits loop
		// } else {
		// 	fmt.Println("Invalid entry")
		// }

		// Logarithmic Search Strategy

		if response == "a" {
			high = guess - 1
		} else if response == "b" {
			low = guess + 1
		} else if response == "c" {
			fmt.Println("I won!")
			break // exits loop
		} else {
			fmt.Println("Invalid entry")
		}
	}
}

// Stretch
// 1. Print out how many tries it took to win
// 2. Generate a number to guess and play with itself
