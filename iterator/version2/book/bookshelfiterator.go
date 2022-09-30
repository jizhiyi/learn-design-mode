package book

// BookShelfIterator 书架的迭代器
type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func NewBookShelfIterator(bookShelf *BookShelf) *BookShelfIterator {
	return &BookShelfIterator{bookShelf: bookShelf}
}

// HasNext 实现 iterator.Iterator
func (bI *BookShelfIterator) HasNext() bool {
	return bI.index < bI.bookShelf.GetLength()
}

// Next 实现 iterator.Iterator
func (bI *BookShelfIterator) Next() *Book {
	book := bI.bookShelf.GetBookAt(bI.index)
	bI.index++
	return book
}
