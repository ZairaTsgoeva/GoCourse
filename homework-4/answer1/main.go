package main

import (
	"GoCourse/homework-4/answer1/book"
	"GoCourse/homework-4/answer1/readable"
	"GoCourse/homework-4/answer1/tablet"
)

func main()  {
	objects := []readable.Readable{
		tablet.CreateTablet(100),
		tablet.CreateTablet(2.5),
		book.CreateBook(0),
		book.CreateBook(99),
	}
	read(objects)
	read(objects)
}
func read(objects []readable.Readable)  {
	for _, r := range objects {
		println(r.State())
		if r.Open() {
			println(r.State())
			r.Read()
		}
		println(r.State())
	}
}