package sample

type NoSupport struct {
	*supportBase
}

func NewNoSupport(name string) *NoSupport {
	n := &NoSupport{}
	n.supportBase = &supportBase{
		Support: n,
		Name:    name,
	}
	return n
}

func (o *NoSupport) Resolve(*Trouble) bool {
	return false
}
