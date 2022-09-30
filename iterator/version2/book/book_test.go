package book

import (
	"testing"
)

func TestBook(t *testing.T) {
	bookShelf := NewBookShelf(4)
	bookShelf.AppendBook(NewBook("A Book"))
	bookShelf.AppendBook(NewBook("B Book"))
	bookShelf.AppendBook(NewBook("C Book"))
	bookShelf.AppendBook(NewBook("D Book"))
	it := bookShelf.Iterator()
	for it.HasNext() {
		b := it.Next()
		t.Log(b.GetName())
	}
}
