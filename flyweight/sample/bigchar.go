package sample

import (
	"fmt"
	"os"
)

type BigChar struct {
	charName byte
	fontData string
}

func NewBigChar(ch byte) *BigChar {
	b := &BigChar{charName: ch}
	fileName := fmt.Sprintf("big%c.txt", ch)
	buf, err := os.ReadFile(fileName)
	if err != nil {
		b.fontData = fmt.Sprintf("%c?", ch)
	} else {
		b.fontData = string(buf)
	}
	return b
}

func (b *BigChar) print() {
	fmt.Println(b.fontData)
}
