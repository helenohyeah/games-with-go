package main

import "fmt"

// Structs allow you to make your own types
// Allows you to package multiple data types into one entity

type Position struct {
	x float32
	y float32
}

type badGuy struct {
	name   string
	health int
	pos    Position // you can have a struct within a struct
}

func whereIsBadGuy(b badGuy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Println("(", x, ",", y, ")")
}

func main() {
	// var p Position

	// fmt.Println(p.x) // 0
	// fmt.Println(p.y) // 0

	// Struct intialization
	p := Position{4, 2}
	fmt.Println(p.x) // 4

	b := badGuy{"Ted", 100, p}
	fmt.Println(b)
	whereIsBadGuy(b)
}
