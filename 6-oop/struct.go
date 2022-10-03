package main

import "fmt"

type Book struct {
	title  string
	author string
}

func main() {
	var book Book
	book.title = "a"
	book.author = "arther"

	changeBook(book)
	fmt.Printf("book： %v\n", book)

	changeBook2(&book)
	fmt.Printf("book： %v\n", book)
}

func changeBook2(b *Book) {
	b.author = "alice"
}

func changeBook(book Book) {
	book.author = "alice"
}
