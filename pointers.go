package main

import "fmt"

/**
* Easy because of my previous experience with
* C and PHP
*/
func main() {
	book := book{}
	bookCopy := book
	bookCopyByReference := &book
	
	fmt.Println(book.title)
	
	book.updateTitle("The Green Book")
	
	fmt.Println(book.title)
	fmt.Println(bookCopy.title)
	fmt.Println((*bookCopyByReference).title)
}

type book struct {
	title string
	releaseYear int
}

func (b *book) updateTitle(newTitle string) {
	b.title = newTitle
}
