package sample1

import (
	"fmt"
)

type StringDisplay struct {
	str string
}

func NewStringDisplay(str string) *AbstractDisplay {
	ds := &StringDisplay{str: str}
	return &AbstractDisplay{AbstractDisplayIn: ds}
}

func (sd *StringDisplay) open() {
	sd.printLine()
}

func (sd *StringDisplay) print() {
	fmt.Printf("|%s|\n", sd.str)
}

func (sd *StringDisplay) close() {
	sd.printLine()
}

func (sd *StringDisplay) printLine() {
	fmt.Printf("+")
	for i, width := 0, len(sd.str); i < width; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("+\n")
}
