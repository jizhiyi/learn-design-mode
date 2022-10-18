package sample

type StringDisplay struct {
	str string
}

func NewStringDisplay(str string) *StringDisplay {
	return &StringDisplay{str: str}
}

func (s *StringDisplay) GetColumns() int {
	return len(s.str)
}

func (s *StringDisplay) GetRows() int {
	return 1
}

func (s *StringDisplay) GetRowText(row int) string {
	if row == 0 {
		return s.str
	}
	return ""
}
