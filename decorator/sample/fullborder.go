package sample

import "strings"

type FullBorder struct {
	*Border
}

func NewFullBorder(display displayIn) *FullBorder {
	return &FullBorder{Border: &Border{displayIn: display}}
}

func (f *FullBorder) GetColumns() int {
	return 1 + f.Border.GetColumns() + 1
}

func (f *FullBorder) GetRows() int {
	return 1 + f.Border.GetRows() + 1
}

func (f *FullBorder) GetRowText(row int) string {
	if row == 0 {
		return "+" + f.makeLine('-', f.Border.GetColumns()) + "+"
	} else if row == f.Border.GetRows()+1 {
		return "+" + f.makeLine('-', f.Border.GetColumns()) + "+"
	}
	return "|" + f.Border.GetRowText(row-1) + "|"
}

func (f *FullBorder) makeLine(ch byte, count int) string {
	buf := strings.Builder{}
	for i := 0; i < count; i++ {
		buf.WriteByte(ch)
	}
	return buf.String()
}
