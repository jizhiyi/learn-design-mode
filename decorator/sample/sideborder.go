package sample

import "fmt"

type SideBorder struct {
	*Border
	sidech byte
}

func NewSideBorder(display displayIn, sidech byte) *SideBorder {
	return &SideBorder{Border: &Border{displayIn: display}, sidech: sidech}
}

func (s *SideBorder) GetColumns() int {
	return 1 + s.Border.GetColumns() + 1
}

func (s *SideBorder) GetRows() int {
	return s.Border.GetRows()
}

func (s *SideBorder) GetRowText(row int) string {
	return fmt.Sprintf("%c%s%c", s.sidech, s.Border.GetRowText(row), s.sidech)
}
