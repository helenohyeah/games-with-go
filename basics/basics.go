package main // tells go this is the file to run as the executable versus a library that can be imported

import "fmt"

func main() {

	// variables()

	loops()
}

func variables() {
	// Declaring Variables

	var a int // if you don't set a value, go will assign a zero value
	var b string
	var c bool
	fmt.Println(a) // 0
	fmt.Println(b) // empty string
	fmt.Println(c) // false

	// Integers

	// var a int8 // 8bits -127, 128
	// var b uint8 // unsigned meaning numbers can only be positive 0, 255
	// var c int16 // also 32, 64
	// var d uint16
	// var c int // 32bits on 32bit machines, 64bits on 64bit machines

	// Floats

	// var f1 float32 // eg 3.14159
	// var f2 float64 // 64bits so you can represent larger or more precise floats

	// Types in Go

	// you can't mismatch types in go
	// you must explicitly tell go to convert types
	i := 5
	f := 3.14159
	// fmt.Println(a + b) // error invalid operation: mismatched types
	fmt.Println(i + int(f))     // 8 float gets truncated, not rounded
	fmt.Println(float64(i) + f) // 8.14159

	// Math is not Real
	// there are libraries that can handle precise math. The more precise, the slower it tends to be

	var x uint8 = 255
	var y uint8 = 1
	var z uint8 = 10

	// go will wrap around if result exceeds number that can be expressed in bits
	fmt.Println(x + y) // 0
	fmt.Println(x + z) // 9

	// be careful using floats because computer has limited bits so answer is always approximate
	// over time these small errors can build up
	f1 := 2.0
	f2 := 3.0
	fmt.Printf("%.20f\n", f1/f2) // 0.66666666666666662966 %.20f means print out all decimal points

	// Booleans

	true := true // stored in 8 bits
	false := false
	fmt.Println(true)  // true
	fmt.Println(false) // false

	// Strings
	// use double quotes

	helloWorld := "Hello World"
	fmt.Println(helloWorld)
}

func loops() {
	// go only has the for loop

	// Infinite Loop

	// for {
	// 	fmt.Println("Hello World")
	// }

	// While Loop

	// i := 0
	// for i < 10 {
	// 	fmt.Println("Hello World", i)
	// 	i = i + 1
	// }

	// For Loop

	for i := 0; i < 10; i++ {
		fmt.Println("Hello World", i)
	}

	// Conditional Loop

	true := true
	if true {
		fmt.Println("Hello World")
	} else {
		fmt.Println("Goodbye")
	}
	// or
	if !true {
		fmt.Println("Goodbye")
	}

	x := 5
	if x > 5 {
		fmt.Println("Less than 5")
	} else if x < 5 {
		fmt.Println("Greater than 5")
	} else {
		fmt.Println("It is 5")
	}
}
