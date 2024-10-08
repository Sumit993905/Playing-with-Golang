package main

import "fmt"

type Book struct {
	name    string
	edition int
	writer  string
}

func main() {

	Book1 := Book{
		name:    "How to win friends and influence people",
		edition: 100,
		writer:  "Dale Carnegi",
	}
	Book1.edition = 100
	fmt.Println(Book1.name)
	fmt.Println(Book1.edition)

}
