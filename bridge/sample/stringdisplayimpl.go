package sample

import "fmt"

type StringDisplayImpl struct {
	str string
}

func NewStringDisplayImpl(str string) *StringDisplayImpl {
	return &StringDisplayImpl{str: str}
}

func (s *StringDisplayImpl) RawOpen() {
	s.printline()
}

func (s *StringDisplayImpl) RawPrint() {
	fmt.Printf("|%s|", s.str)
}

func (s *StringDisplayImpl) RawClose() {
	s.printline()
}

func (s *StringDisplayImpl) printline() {
	fmt.Println("+")
	for i, sz := 0, len(s.str); i < sz; i++ {
		fmt.Println("-")
	}
	fmt.Println("+")
}
