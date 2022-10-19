package sample

import "fmt"

type ListVisitor struct {
	currentdir string
}

func (l *ListVisitor) Visit(entry Entry) {
	fmt.Printf("%s/%s\n", l.currentdir, entry)
	savedir := l.currentdir
	l.currentdir = l.currentdir + "/" + entry.GetName()
	iter := entry.Iterator()
	for iter.HasNext() {
		l.Visit(iter.Next())
	}
	l.currentdir = savedir
}
