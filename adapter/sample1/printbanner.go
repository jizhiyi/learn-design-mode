package sample1

type PrintBanner struct {
	*Banner
}

func NewPrintBanner(cont string) *PrintBanner {
	return &PrintBanner{Banner: &Banner{cont: cont}}
}

func (p *PrintBanner) PrintWeak() {
	p.Banner.showWithParen()
}
func (p *PrintBanner) PrintStrong() {
	p.Banner.showWithAster()
}
