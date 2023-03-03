package iterator

type IterableCollection interface {
	createIterator() iterator
}

type iterator interface {
	hasNext() bool
	next() *Book
}

type BookIterator struct {
	current int
	books []Book
}


func (bi *BookIterator) hasNext() bool {
	return bi.current < len(bi.books)
}

func (bi *BookIterator) next() *Book {
	if bi.hasNext() {
		bk := bi.books[bi.current]
		bi.current++
		return &bk
	}
	return nil
}