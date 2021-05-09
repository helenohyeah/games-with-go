package main

import "fmt"

// Pointers are always working behind the scenes

// Many languages hide it from you (e.g. Java)
// Some languages let you modify them (e.g. C)
// Some languages fall in the middle (e.g. C#, Go)

// Go gives you visibility into pointers but doesn't let you modify them

type Position struct {
	x float32
	y float32
}

type badGuy struct {
	name   string
	health int
	pos    Position // you can have a struct within a struct
}

func whereIsBadGuy(b *badGuy) {
	// Go will automatically dereference members in a struct when you work with a pointer to a struct
	x := b.pos.x // Do not need to dereference to access
	y := b.pos.y
	fmt.Println("(", x, ",", y, ")")
}

func addOne(x int) {
	x = x + 1
}

func addOneWithPointer(x *int) {
	// x = x + 1 // error - we cannot add 1 to a pointer
	// We have to dereference the pointer
	*x = *x + 1 // *x = give me the thing that the pointer points to
}

func main() {

	// Values
	x := 5
	fmt.Println(x) // 5

	// Pointer
	// &  = give me the address in memory of that value
	var xPointer *int = &x // type *int
	fmt.Println(xPointer)  // 0xc000118000 - hexadecimal virtual memory address

	addOne(x)
	fmt.Println(x) // 5 - a copy is modified in addOne function and the original is unchanged

	// Use case: Change the thing you pass it rather than make a copy and change the copy
	addOneWithPointer(xPointer)
	fmt.Println(x) // 6

	// Pass the struct around instead of making many copies
	p := Position{4, 2}
	b := badGuy{"Ted", 100, p}
	whereIsBadGuy(&b)
}
