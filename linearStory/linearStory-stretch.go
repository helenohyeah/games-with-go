package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

// Stretch
// 1. Add a function that will insert a new page, after a given page
// 2. Add a function that will delete a page
// 3. Load a story from a file

// Example of a linked list data structure
type storyPage struct {
	text     string
	nextPage *storyPage
}

func getInput(scanner *bufio.Scanner) string {
	fmt.Print("> ")
	scanner.Scan()
	return scanner.Text()
}

func playStory(p *storyPage) {
	if p == nil {
		return
	}

	fmt.Println(p.text)
	playStory(p.nextPage)
}

func addPage(p *storyPage) {
	scanner := bufio.NewScanner(os.Stdin)

	if p == nil {
		return
	}

	for {
		fmt.Println("Page:", p.text)
		fmt.Println("Insert a new page after this one? [y/n]")

		input := getInput(scanner)

		if input == "y" {
			fmt.Println("Enter text for new page:")
			text := getInput(scanner)

			newPage := storyPage{
				text:     text,
				nextPage: p.nextPage,
			}
			p.nextPage = &newPage
			fmt.Println("===> PAGE ADDED")
			return
		} else if input == "n" {
			addPage(p.nextPage)
			return
		} else {
			fmt.Println("x==> INVALID ENTRY")
		}
	}
}

func removePage(p *storyPage, firstPage storyPage, prevPage *storyPage) storyPage {
	scanner := bufio.NewScanner(os.Stdin)

	if p == nil {
		return firstPage
	}

	for {
		fmt.Println("Page:", p.text)
		fmt.Println("Remove this page? [y/n]")

		input := getInput(scanner)

		if input == "y" {
			// Modify story's first page if it is being removed
			if reflect.DeepEqual(*p, firstPage) {
				if p.nextPage == nil {
					fmt.Println("x==> CANNOT REMOVE LAST PAGE")
					return firstPage
				}
				fmt.Println("===> PAGE REMOVED")
				return *p.nextPage
			}
			// Modify pointer to previous page
			if p.nextPage == nil {
				*prevPage = storyPage{
					text:     prevPage.text,
					nextPage: nil,
				}
			} else {
				*prevPage.nextPage = *p.nextPage
			}
			fmt.Println("===> PAGE REMOVED")
			return firstPage
		} else if input == "n" {
			removePage(p.nextPage, firstPage, p)
			return firstPage
		} else {
			fmt.Println("x==> INVALID ENTRY")
		}
	}
}

func getNextPage(page *storyPage) *storyPage {
	if page.nextPage != nil {
		return page.nextPage
	}

	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	p1 := storyPage{"It was a dark and stormy night.", nil}
	p2 := storyPage{"You are alone and you need to find the sacred nail before the evil doers do.", nil}
	p3 := storyPage{"You see a troll ahead.", nil}
	p1.nextPage = &p2
	p2.nextPage = &p3

	fmt.Println("===> BEGIN")

	for {
		if &p1 == nil {
			fmt.Println("x==> CANNOT DELETE LAST PAGE IN STORY")
			break
		}

		fmt.Println("===> CHOOSE AN OPTION")
		fmt.Println("Play story     [p]")
		fmt.Println("Add a page     [a]")
		fmt.Println("Remove a page  [r]")
		fmt.Println("Exit           [e]")

		input := getInput(scanner)

		if input == "p" {
			fmt.Println("===> PLAYING STORY...")
			playStory(&p1)
		} else if input == "a" {
			fmt.Println("===> ADDING PAGE...")
			addPage(&p1)
		} else if input == "r" {
			fmt.Println("===> REMOVING PAGE...")
			newFirstPage := removePage(&p1, p1, nil)
			p1 = newFirstPage
		} else if input == "e" {
			break
		} else {
			fmt.Println("x==> INVALID ENTRY")
		}
	}
}
