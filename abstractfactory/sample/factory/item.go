package factory

type Item interface {
	MakeHTML() string
}

type ItemSt struct {
	Caption string
}
