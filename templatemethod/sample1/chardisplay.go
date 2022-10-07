package sample1

import "fmt"

type CharDisplay struct {
	ch byte
}

func NewCharDisplay(ch byte) *AbstractDisplay {
	cd := &CharDisplay{ch: ch}
	return &AbstractDisplay{cd}
}

func (cd *CharDisplay) open() {
	fmt.Printf("<<")
}

func (cd *CharDisplay) print() {
	fmt.Printf("%c", cd.ch)
}

func (cd *CharDisplay) close() {
	fmt.Printf(">>\n")
}
