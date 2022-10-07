package sample2

import "fmt"

type StringDisplay struct {
	*AbstractDisplay
	str string
}

func NewStringDisplay(str string) *StringDisplay {
	sd := &StringDisplay{str: str}
	sd.AbstractDisplay = &AbstractDisplay{AbstractDisplayIn: sd}
	return sd
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
