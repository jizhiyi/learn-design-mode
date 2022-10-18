package sample

import "fmt"

type Entry interface {
	GetName() string
	GetSize() int
	Add(entry Entry)
	PrintList()
	printList(prefix string)
	String() string
}

type entryBase struct {
	Entry
	Name string
}

func (e *entryBase) GetName() string {
	return e.Name
}

func (e *entryBase) Add(entry Entry) {
	panic("son not need add")
}

func (e *entryBase) String() string {
	return fmt.Sprintf("%s(%d)", e.GetName(), e.GetSize())
}

func (d *entryBase) PrintList() {
	d.printList("")
}
