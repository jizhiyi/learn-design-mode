package sample

type OddSupport struct {
	*supportBase
}

func NewOddSupport(name string) *OddSupport {
	o := &OddSupport{}
	o.supportBase = &supportBase{
		Support: o,
		Name:    name,
	}
	return o
}

func (o *OddSupport) Resolve(trouble *Trouble) bool {
	return trouble.GetNumber()%2 == 1
}
