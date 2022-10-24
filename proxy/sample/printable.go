package sample

type Printable interface {
	SetPrintName(name string)
	GetPrintName() string
	Print(str string)
}
