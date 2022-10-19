package sample

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

func (d *Directory) Iterator() EntryIterator {
	return &drectoryIterator{index: 0, directory: d}
}
