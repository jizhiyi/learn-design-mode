package sample

type PrinterProxy struct {
	name string
	real *Printer
}

func NewPrinterProxy(name string) *PrinterProxy {
	return &PrinterProxy{name: name}
}

func (p *PrinterProxy) SetPrintName(name string) {
	if p.real != nil {
		p.real.SetPrintName(name)
	}
	p.name = name
}

func (p *PrinterProxy) GetPrintName() string {
	return p.name
}

func (p *PrinterProxy) Print(str string) {
	p.realize()
	p.real.Print(str)
}

func (p *PrinterProxy) realize() {
	if p.real == nil {
		p.real = NewPrinter(p.name)
	}
}
