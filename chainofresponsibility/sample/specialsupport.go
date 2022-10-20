package sample

type SpecialSupport struct {
	*supportBase
	number int
}

func NewSpecialSupport(name string, number int) *SpecialSupport {
	s := &SpecialSupport{number: number}
	s.supportBase = &supportBase{
		Support: s,
		Name:    name,
	}
	return s
}

func (o *SpecialSupport) Resolve(trouble *Trouble) bool {
	return trouble.GetNumber() == o.number
}
