package sample

import "fmt"

type BigChar struct {
	charName byte
	fontData string
}

func NewBigChar(ch byte) *BigChar {
	return nil
}

func (b *BigChar) print() {
	fmt.Println(b.fontData)
}
