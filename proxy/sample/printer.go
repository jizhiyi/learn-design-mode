package sample

import (
	"fmt"
	"time"
)

type Printer struct {
	name string
}

func NewPrinter(name string) *Printer {
	return &Printer{name: name}
}

func (p *Printer) SetPrintName(name string) {
	p.name = name
}

func (p *Printer) GetPrintName() string {
	return p.name
}

func (p *Printer) Print(str string) {
	fmt.Printf("===%s===\n", p.name)
	fmt.Println(str)
}

func (p *Printer) heavyJob(msg string) {
	fmt.Print(msg)
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println("结束")
}
