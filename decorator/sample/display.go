package sample

import "fmt"

type displayIn interface {
	GetColumns() int
	GetRows() int
	GetRowText(row int) string
}

type Display struct {
	Inner displayIn
}

func NewDisplay(inner displayIn) *Display {
	return &Display{Inner: inner}
}

func (d *Display) Show() {
	row := d.Inner.GetRows()
	for i := 0; i < row; i++ {
		fmt.Println(d.Inner.GetRowText(i))
	}
}
