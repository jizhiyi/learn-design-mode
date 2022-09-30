package book

// Book 书类
type Book struct {
	name string
}

func NewBook(name string) *Book {
	return &Book{name: name}
}

func (book *Book) GetName() string {
	return book.name
}
