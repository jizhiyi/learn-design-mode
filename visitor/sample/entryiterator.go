package sample

type EntryAggregate interface {
	Iterator() EntryIterator
}

type EntryIterator interface {
	HasNext() bool
	Next() Entry
}

var emptyIterator EntryIterator = &emptyEntryIterator{}

type emptyEntryIterator struct {
}

func (*emptyEntryIterator) HasNext() bool {
	return false
}

func (*emptyEntryIterator) Next() Entry {
	return nil
}

type drectoryIterator struct {
	index     int
	directory *Directory
}

func (iter *drectoryIterator) HasNext() bool {
	return iter.index < len(iter.directory.entryList)
}

func (iter *drectoryIterator) Next() Entry {
	if iter.HasNext() {
		entry := iter.directory.entryList[iter.index]
		iter.index++
		return entry
	}
	return nil
}
