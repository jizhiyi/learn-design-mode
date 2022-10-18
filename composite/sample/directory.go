package sample

import "fmt"

type Directory struct {
	*entryBase
	entryList []Entry
}

func NewDirectory(name string) *Directory {
	d := &Directory{}
	base := &entryBase{
		Entry: d,
		Name:  name,
	}
	d.entryBase = base
	return d
}

func (d *Directory) GetSize() (size int) {
	for _, entry := range d.entryList {
		size += entry.GetSize()
	}
	return
}

func (d *Directory) Add(entry Entry) {
	d.entryList = append(d.entryList, entry)
}

func (d *Directory) printList(prefix string) {
	fmt.Printf("%s/%s\n", prefix, d)
	newPrefix := fmt.Sprintf("%s/%s", prefix, d.GetName())
	for _, entry := range d.entryList {
		entry.printList(newPrefix)
	}
}
