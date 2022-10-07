package sample2

import "fmt"

type CharDisplay struct {
	*AbstractDisplay
	ch byte
}

func NewCharDisplay(ch byte) *CharDisplay {
	cd := &CharDisplay{ch: ch}
	cd.AbstractDisplay = NewAbstractDisplay(cd)
	return cd
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

//可以去覆盖
//func (cd *CharDisplay) Display() {
//	cd.open()
//	for i := 0; i < 3; i++ {
//		cd.print()
//	}
//	cd.close()
//}
