package sample

type LimitSupport struct {
	*supportBase
	limit int
}

func NewLimitSupport(name string, limit int) *LimitSupport {
	l := &LimitSupport{limit: limit}
	l.supportBase = &supportBase{
		Support: l,
		Name:    name,
	}
	return l
}

func (l *LimitSupport) Resolve(trouble *Trouble) bool {
	return trouble.GetNumber() < l.limit
}
