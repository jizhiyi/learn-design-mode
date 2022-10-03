package sample2

type PrintBanner struct {
	*Banner
	*Print
}

func NewPrintBanner(cont string) *PrintBanner {
	return &PrintBanner{Banner: &Banner{cont: cont}, Print: &Print{}}
}

func (p *PrintBanner) PrintWeak() {
	p.Banner.showWithParen()
}
func (p *PrintBanner) PrintStrong() {
	p.Banner.showWithAster()
}
