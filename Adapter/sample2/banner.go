package sample2

import "fmt"

type Banner struct {
	cont string
}

func (b *Banner) showWithParen() {
	fmt.Println("(", b.cont, ")")
}

func (b *Banner) showWithAster() {
	fmt.Println("*" + b.cont + "*")
}
