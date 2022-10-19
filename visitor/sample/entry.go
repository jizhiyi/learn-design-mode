package sample

import "fmt"

type Element interface {
	Accept(v Visitor)
}

type Entry interface {
	GetName() string
	GetSize() int
	Add(entry Entry)
	String() string
	EntryAggregate
	Element
}

type entryBase struct {
	Entry
	Name string
}

func (e *entryBase) GetName() string {
	return e.Name
}

func (e *entryBase) Add(Entry) {
	panic("son not need add")
}

func (e *entryBase) String() string {
	return fmt.Sprintf("%s(%d)", e.GetName(), e.GetSize())
}

func (e *entryBase) Iterator() EntryIterator {
	return emptyIterator
}

func (e *entryBase) Accept(v Visitor) {
	v.Visit(e.Entry)
}
