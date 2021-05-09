package main

import "fmt"

// Using functions to DRY up your code

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}

func beSociable(name string) {
	sayHello(name)
	fmt.Println("How's the weather?")
	sayGoodbye(name)
}

// Using functions to return values

func addOne(x int) int {
	return x + 1
}

// Recursive function

func sayHelloABunch() {
	fmt.Println("Hello")
	sayHelloABunch()
}

func main() {
	beSociable("Helen")
	beSociable("Bob")

	x := 5
	x = addOne(x)
	fmt.Println(x)

	// Compose functions: take the result of a function and directly pass it into another function
	x = addOne(addOne(x)) // 8
	fmt.Println(x)

	sayHelloABunch()
}
