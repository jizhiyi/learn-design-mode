package sample

type File struct {
	*entryBase
	size int
}

func NewFile(name string, size int) *File {
	f := &File{size: size}
	base := &entryBase{
		Entry: f,
		Name:  name,
	}
	f.entryBase = base
	return f
}

func (f *File) GetSize() int {
	return f.size
}
