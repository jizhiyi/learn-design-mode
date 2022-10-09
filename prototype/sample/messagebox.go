package sample

import (
	"fmt"
	"learn/prototype/sample/framework"
)

type MessageBox struct {
	decocher byte
}

func NewMessageBox(decocher byte) *MessageBox {
	return &MessageBox{decocher: decocher}
}

// 功能方法
func (m *MessageBox) Use(str string) {
	sz := len(str)
	for i := 0; i < sz+4; i++ {
		fmt.Printf("%c", m.decocher)
	}
	fmt.Println()
	fmt.Printf("%c %s %c\n", m.decocher, str, m.decocher)
	for i := 0; i < sz+4; i++ {
		fmt.Printf("%c", m.decocher)
	}
	fmt.Println()
}

// 复制方法
func (m *MessageBox) Clone() framework.Product {
	return &MessageBox{decocher: m.decocher}
}
