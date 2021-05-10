package main

import (
	"fmt"
)

// Example of a linked list data structure
type storyPage struct {
	text string
	// nextPage storyPage // error => memory leak
	nextPage *storyPage
}

func playStory(p *storyPage) {
	// Base case
	if p == nil {
		return
	}

	// Without nil pointer check, will result in runtime error
	fmt.Println(p.text) // error => invalid memory address or nil pointer dereference
	playStory(p.nextPage)
}

func main() {
	// scanner := bufio.NewScanner(os.Stdin)

	// fmt.Println("It was a dark and story night.")
	// scanner.Scan()
	// fmt.Println("You are alone and you need to find the sacred nail before the evil doers do.")
	// scanner.Scan()
	// fmt.Println("You see a troll ahead.")

	p1 := storyPage{"It was a dark and story night.", nil}
	p2 := storyPage{"You are alone and you need to find the sacred nail before the evil doers do.", nil}
	p3 := storyPage{"You see a troll ahead.", nil}
	p1.nextPage = &p2
	p2.nextPage = &p3

	playStory(&p1)
}

// Stretch
// 1. Add a function that will insert a new page, after a given page
// 2. Add a function that will delete a page
// 3. Load a story from a file
