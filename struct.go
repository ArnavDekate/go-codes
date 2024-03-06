package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var book1 Books
	var book2 Books

	book1.title = "Go Programming"
	book1.author = "Madan Tiwari"
	book1.subject = "Go Programming Tutorial"
	book1.book_id = 637383

	book2.title = "Telecom Billing"
	book2.author = "Rajesh"
	book2.subject = "Telecom billing tutorial"
	book2.book_id = 633422

	fmt.Printf("Book1 Title : %s\n", book1.title)
	fmt.Printf("Book1 Author : %s\n", book1.author)
	fmt.Printf("Book1 Subject : %s\n", book1.subject)
	fmt.Printf("Book1 Id : %d\n\n", book1.book_id)

	fmt.Printf("Book2 Title : %s\n", book2.title)
	fmt.Printf("Book2 Author : %s\n", book2.author)
	fmt.Printf("Book2 Subject : %s\n", book2.subject)
	fmt.Printf("Book2 Id : %d\n\n", book2.book_id)
	fmt.Println(book1, book2)

	book3 := Books{
		title:   "C Program",
		author:  "Arnav",
		subject: "C Tutorial",
		book_id: 635473}
	prtbook(book3)

}

func prtbook(book Books) {
	fmt.Printf("Book2 Title : %s\n", book.title)
	fmt.Printf("Book2 Author : %s\n", book.author)
	fmt.Printf("Book2 Subject : %s\n", book.subject)
	fmt.Printf("Book2 Id : %d\n\n", book.book_id)
}
