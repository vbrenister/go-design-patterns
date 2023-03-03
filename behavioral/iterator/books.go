package iterator

import "fmt"


type BookType int

const (
	HardCover BookType = iota
	SoftCover
	PaperBack
	EBook
)

type Book struct {
	Name string 
	Author string
	PageCount int
	Type BookType
}

type Library struct {
	Collection []Book
}

func (l *Library) createIterator() iterator {
	return &BookIterator{
		books: l.Collection,
	}
}

func (l *Library) IterateBooks(f func (Book) error)  {
	for _, book := range l.Collection {
		err := f(book)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}