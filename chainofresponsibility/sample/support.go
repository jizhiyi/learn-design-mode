package sample

import "fmt"

type Support interface {
	SetNext(next Support) Support
	SupportFun(trouble *Trouble)
	Resolve(trouble *Trouble) bool
}

type supportBase struct {
	Support
	Name        string
	nextSupport Support
}

func (s *supportBase) SetNext(next Support) Support {
	s.nextSupport = next
	return next
}

func (s *supportBase) SupportFun(trouble *Trouble) {
	if s.Resolve(trouble) {
		s.done(trouble)
	} else if s.nextSupport != nil {
		s.nextSupport.SupportFun(trouble)
	} else {
		s.fail(trouble)
	}
}

func (s *supportBase) String() string {
	return fmt.Sprintf("[%s]", s.Name)
}

func (s *supportBase) done(trouble *Trouble) {
	fmt.Printf("%s is resolved by %s.\n", trouble, s)
}

func (s *supportBase) fail(trouble *Trouble) {
	fmt.Printf("%s cannot be resolved.", trouble)
}
