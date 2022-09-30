package book

import iterator "learn/iterator/version2"

// BookShelf 书架类
type BookShelf struct {
	books []*Book
}

func NewBookShelf(initCap int) *BookShelf {
	return &BookShelf{books: make([]*Book, 0, initCap)}
}

func (bs *BookShelf) GetBookAt(index int) *Book {
	return bs.books[index]
}

func (bs *BookShelf) AppendBook(book *Book) {
	bs.books = append(bs.books, book)
}

func (bs *BookShelf) GetLength() int {
	return len(bs.books)
}

// Iterator 实现 iterator.Aggregate
func (bs *BookShelf) Iterator() iterator.Iterator[*Book] {
	return NewBookShelfIterator(bs)
}
