package sample2

type Print struct {
}

func (p *Print) PrintWeak() {
	panic("not implements")
}

func (p *Print) PrintStrong() {
	panic("not implements")
}
