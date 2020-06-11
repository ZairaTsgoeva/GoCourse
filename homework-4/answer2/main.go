package main

import (
	"GoCourse/homework-4/answer2/contact"
	"sort"
)

func main()  {
	book := createBook()

	println("before sort")
	for _, c := range *book {
		println(c.GetName())
	}

	sort.Sort(book)

	println("after sort")
	for _, c := range *book {
		println(c.GetName())
	}
}

func createBook() *contact.Book {
	mask := contact.CreateContact("Маск")
	ivan := contact.CreateContact("Иван")
	mask.AddPhone(02)
	mask.AddPhone(03)
	ivan.AddPhone(04)
	ivan.AddPhone(05)

	book := contact.CreateBook()
	book.AddContact(mask)
	book.AddContact(ivan)

	return book
}
