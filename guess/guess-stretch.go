package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Stretch
// 1. Print out how many tries it took to win
// 2. Generate a number to guess and play with itself

// A better number guessing program
func main() {
	// butfio reads things from input sources
	// os.Stdin means input source from console
	scanner := bufio.NewScanner((os.Stdin))

	low := 1
	high := 100

	// number := rand.Intn(high) // number generator is deterministic so it will produce the same number each time
	// use a seed that changes to produce varying random numbers
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	number := r.Intn(high)

	tries := 0

	fmt.Printf("I'm going to think of a number between %v and %v", low, high)
	fmt.Printf("\nMy number is: %v", number)
	fmt.Println("\nPrint ENTER when ready")
	scanner.Scan() // wait for input

	// Binary Search Strategy
	// guess := (low + high) / 2
	// for {
	// 	fmt.Printf("\nMy guess: %v\n", guess)

	// 	if guess < number {
	// 		fmt.Println("Too low")
	// 		guess++
	// 	} else if guess > number {
	// 		fmt.Println("Too high")
	// 		guess--
	// 	} else {
	// 		fmt.Println("Correct!")
	// 		fmt.Printf("\nYou took %v tries to win\n", tries)
	// 		break
	// 	}

	// 	tries++
	// }

	// Logarithmic Search Strategy
	for {
		guess := (low + high) / 2
		fmt.Printf("\nMy guess: %v\n", guess)

		if guess < number {
			fmt.Println("Too low")
			low = guess + 1
		} else if guess > number {
			fmt.Println("Too high")
			high = guess - 1
		} else {
			fmt.Println("Correct!")
			fmt.Printf("\nYou took %v tries to win\n", tries)
			break
		}

		tries++
	}
}
